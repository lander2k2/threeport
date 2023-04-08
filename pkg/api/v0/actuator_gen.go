// generated by 'threeport-codegen api-model' - do not edit

package v0

import (
	"encoding/json"
	"fmt"
	notifications "github.com/threeport/threeport/pkg/notifications"
)

const (
	ObjectTypeProfile ObjectType = "Profile"
	ObjectTypeTier    ObjectType = "Tier"

	ActuatorStreamName = "actuatorStream"

	ProfileSubject       = "profile.*"
	ProfileCreateSubject = "profile.create"
	ProfileUpdateSubject = "profile.update"
	ProfileDeleteSubject = "profile.delete"

	TierSubject       = "tier.*"
	TierCreateSubject = "tier.create"
	TierUpdateSubject = "tier.update"
	TierDeleteSubject = "tier.delete"

	PathProfiles = "/v0/profiles"
	PathTiers    = "/v0/tiers"
)

// GetProfileSubjects returns the NATS subjects
// for profiles.
func GetProfileSubjects() []string {
	return []string{
		ProfileCreateSubject,
		ProfileUpdateSubject,
		ProfileDeleteSubject,
	}
}

// GetTierSubjects returns the NATS subjects
// for tiers.
func GetTierSubjects() []string {
	return []string{
		TierCreateSubject,
		TierUpdateSubject,
		TierDeleteSubject,
	}
}

// GetActuatorSubjects returns the NATS subjects
// for all actuator objects.
func GetActuatorSubjects() []string {
	var actuatorSubjects []string

	actuatorSubjects = append(actuatorSubjects, GetProfileSubjects()...)
	actuatorSubjects = append(actuatorSubjects, GetTierSubjects()...)

	return actuatorSubjects
}

// NotificationPayload returns the notification payload that is delivered to the
// controller when a change is made.  It includes the object as presented by the
// client when the change was made.
func (p *Profile) NotificationPayload(requeue bool, lastDelay int64) (*[]byte, error) {
	notif := notifications.Notification{
		LastRequeueDelay: &lastDelay,
		Object:           p,
		Requeue:          requeue,
	}

	payload, err := json.Marshal(notif)
	if err != nil {
		return &payload, fmt.Errorf("failed to marshal notification payload %+v: %w", p, err)
	}

	return &payload, nil
}

// GetID returns the unique ID for the object.
func (p *Profile) GetID() uint {
	return *p.ID
}

// String returns a string representation of the ojbect.
func (p Profile) String() string {
	return fmt.Sprintf("v0.Profile")
}

// NotificationPayload returns the notification payload that is delivered to the
// controller when a change is made.  It includes the object as presented by the
// client when the change was made.
func (t *Tier) NotificationPayload(requeue bool, lastDelay int64) (*[]byte, error) {
	notif := notifications.Notification{
		LastRequeueDelay: &lastDelay,
		Object:           t,
		Requeue:          requeue,
	}

	payload, err := json.Marshal(notif)
	if err != nil {
		return &payload, fmt.Errorf("failed to marshal notification payload %+v: %w", t, err)
	}

	return &payload, nil
}

// GetID returns the unique ID for the object.
func (t *Tier) GetID() uint {
	return *t.ID
}

// String returns a string representation of the ojbect.
func (t Tier) String() string {
	return fmt.Sprintf("v0.Tier")
}
