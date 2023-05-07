// generated by 'threeport-codegen controller' - do not edit

package workload

import (
	"errors"
	"fmt"
	mapstructure "github.com/mitchellh/mapstructure"
	v0 "github.com/threeport/threeport/pkg/api/v0"
	client "github.com/threeport/threeport/pkg/client/v0"
	controller "github.com/threeport/threeport/pkg/controller"
	notifications "github.com/threeport/threeport/pkg/notifications"
)

// WorkloadInstanceReconciler reconciles system state when a WorkloadInstance
// is created, updated or deleted.
func WorkloadInstanceReconciler(r *controller.Reconciler) {
	r.ShutdownWait.Add(1)
	reconcilerLog := r.Log.WithValues("reconcilerName", r.Name)
	reconcilerLog.Info("reconciler started")
	shutdown := false

	for {
		// create a fresh log object per reconciliation loop so we don't
		// accumulate values across multiple loops
		log := r.Log.WithValues("reconcilerName", r.Name)

		if shutdown {
			break
		}

		// check for shutdown instruction
		select {
		case <-r.Shutdown:
			shutdown = true
		default:
			// pull message off queue
			msg := r.PullMessage()
			if msg == nil {
				continue
			}

			// consume message data to capture notification from API
			notif, err := notifications.ConsumeMessage(msg.Data)
			if err != nil {
				log.Error(
					err, "failed to consume message data from NATS",
					"msgSubject", msg.Subject,
					"msgData", string(msg.Data),
				)
				go r.RequeueRaw(msg.Subject, msg.Data)
				log.V(1).Info("workload instance reconciliation requeued with identical payload and fixed delay")
				continue
			}

			// decode the object that was created
			var workloadInstance v0.WorkloadInstance
			mapstructure.Decode(notif.Object, &workloadInstance)
			log = log.WithValues("workloadInstanceID", workloadInstance.ID)

			// back off the requeue delay as needed
			requeueDelay := controller.SetRequeueDelay(
				notif.LastRequeueDelay,
				controller.DefaultInitialRequeueDelay,
				controller.DefaultMaxRequeueDelay,
			)

			// build the notif payload for requeues
			notifPayload, err := workloadInstance.NotificationPayload(
				notif.Operation,
				true,
				requeueDelay,
			)
			if err != nil {
				log.Error(err, "failed to build notification payload for requeue")
				go r.RequeueRaw(msg.Subject, msg.Data)
				log.V(1).Info("workload instance reconciliation requeued with identical payload and fixed delay")
				continue
			}

			// check for lock on object
			locked, ok := r.CheckLock(&workloadInstance)
			if locked || ok == false {
				go r.Requeue(&workloadInstance, msg.Subject, notifPayload, requeueDelay)
				log.V(1).Info("workload instance reconciliation requeued")
				continue
			}

			// put a lock on the reconciliation of the created object
			if ok := r.Lock(&workloadInstance); !ok {
				go r.Requeue(&workloadInstance, msg.Subject, notifPayload, requeueDelay)
				log.V(1).Info("workload instance reconciliation requeued")
				continue
			}

			// retrieve latest version of object if requeued
			if notif.Requeue {
				latestWorkloadInstance, err := client.GetWorkloadInstanceByID(
					r.APIClient,
					r.APIServer,
					*workloadInstance.ID,
				)
				// check if error is 404 - if object no longer exists, no need to requeue
				if errors.Is(err, client.ErrorObjectNotFound) {
					log.Info(fmt.Sprintf(
						"object with ID %d no longer exists - halting reconciliation",
						*workloadInstance.ID,
					))
					r.ReleaseLock(&workloadInstance)
					continue
				}
				if err != nil {
					log.Error(err, "failed to get workload instance by ID from API")
					r.UnlockAndRequeue(&workloadInstance, msg.Subject, notifPayload, requeueDelay)
					continue
				}
				workloadInstance = *latestWorkloadInstance
			}

			// determine which operation and act accordingly
			switch notif.Operation {
			case notifications.NotificationOperationCreated:
				if err := workloadInstanceCreated(r, &workloadInstance, &log); err != nil {
					log.Error(err, "failed to reconcile created workload instance object")
					r.UnlockAndRequeue(
						&workloadInstance,
						msg.Subject,
						notifPayload,
						requeueDelay,
					)
					continue
				}
			case notifications.NotificationOperationDeleted:
				if err := workloadInstanceDeleted(r, &workloadInstance, &log); err != nil {
					log.Error(err, "failed to reconcile deleted workload instance object")
					r.UnlockAndRequeue(
						&workloadInstance,
						msg.Subject,
						notifPayload,
						requeueDelay,
					)
					continue
				}
			default:
				log.Error(
					errors.New("unrecognized notifcation operation"),
					"notification included an invalid operation",
				)
				r.UnlockAndRequeue(
					&workloadInstance,
					msg.Subject,
					notifPayload,
					requeueDelay,
				)
				continue

			}

			// set the object's Reconciled field to true
			objectReconciled := true
			reconciledWorkloadInstance := v0.WorkloadInstance{
				Common:     v0.Common{ID: workloadInstance.ID},
				Reconciled: &objectReconciled,
			}
			updatedWorkloadInstance, err := client.UpdateWorkloadInstance(
				r.APIClient,
				r.APIServer,
				&reconciledWorkloadInstance,
			)
			if err != nil {
				log.Error(err, "failed to update workload instance to mark as reconciled")
				r.UnlockAndRequeue(&workloadInstance, msg.Subject, notifPayload, requeueDelay)
				continue
			}
			log.V(1).Info(
				"workload instance marked as reconciled in API",
				"workload instanceName", updatedWorkloadInstance.Name,
			)

			// release the lock on the reconciliation of the created object
			if ok := r.ReleaseLock(&workloadInstance); !ok {
				log.V(1).Info("workload instance remains locked - will unlock when TTL expires")
			} else {
				log.V(1).Info("workload instance unlocked")
			}

			log.Info("workload instance successfully reconciled")
		}
	}

	r.Sub.Unsubscribe()
	reconcilerLog.Info("reconciler shutting down")
	r.ShutdownWait.Done()
}
