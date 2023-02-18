// Copyright (c) 2021 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

import (
	"time"
)

type GroupMemberAddMode string

const (
	GroupMemberAddModeAdmin GroupMemberAddMode = "admin_add"
)

// GroupInfo contains basic information about a group chat on WhatsApp.
type GroupInfo struct {
	JID      JID `json:"jid"`
	OwnerJID JID `json:"owner_jid"`

	GroupName
	GroupTopic
	GroupLocked
	GroupAnnounce
	GroupEphemeral

	GroupParent
	GroupLinkedParent
	GroupIsDefaultSub

	GroupCreated time.Time `json:"group_created"`

	ParticipantVersionID string             `json:"participant_version_id"`
	Participants         []GroupParticipant `json:"participants"`

	MemberAddMode GroupMemberAddMode `json:"member_add_mode"`
}

type GroupParent struct {
	IsParent                      bool   `json:"is_parent"`
	DefaultMembershipApprovalMode string `json:"default_membership_approval_mode"` // request_required
}

type GroupLinkedParent struct {
	LinkedParentJID JID `json:"linked_parent_jid"`
}

type GroupIsDefaultSub struct {
	IsDefaultSubGroup bool `json:"is_default_sub_group"`
}

// GroupName contains the name of a group along with metadata of who set it and when.
type GroupName struct {
	Name      string    `json:"name"`
	NameSetAt time.Time `json:"name_set_at"`
	NameSetBy JID       `json:"name_set_by"`
}

// GroupTopic contains the topic (description) of a group along with metadata of who set it and when.
type GroupTopic struct {
	Topic        string    `json:"topic"`
	TopicID      string    `json:"topic_id"`
	TopicSetAt   time.Time `json:"topic_set_at"`
	TopicSetBy   JID       `json:"topic_set_by"`
	TopicDeleted bool      `json:"topic_deleted"`
}

// GroupLocked specifies whether the group info can only be edited by admins.
type GroupLocked struct {
	IsLocked bool `json:"is_locked"`
}

// GroupAnnounce specifies whether only admins can send messages in the group.
type GroupAnnounce struct {
	IsAnnounce        bool   `json:"is_announce"`
	AnnounceVersionID string `json:"announce_version_id"`
}

// GroupParticipant contains info about a participant of a WhatsApp group chat.
type GroupParticipant struct {
	JID          JID  `json:"jid"`
	IsAdmin      bool `json:"is_admin"`
	IsSuperAdmin bool `json:"is_super_admin"`

	// When creating groups, adding some participants may fail.
	// In such cases, the error code will be here.
	Error      int                         `json:"error"`
	AddRequest *GroupParticipantAddRequest `json:"add_request"`
}

type GroupParticipantAddRequest struct {
	Code       string    `json:"code"`
	Expiration time.Time `json:"expiration"`
}

// GroupEphemeral contains the group's disappearing messages settings.
type GroupEphemeral struct {
	IsEphemeral       bool   `json:"is_ephemeral"`
	DisappearingTimer uint32 `json:"disappearing_timer"`
}

type GroupDelete struct {
	Deleted      bool
	DeleteReason string
}

type GroupLinkChangeType string

const (
	GroupLinkChangeTypeParent  GroupLinkChangeType = "parent_group"
	GroupLinkChangeTypeSub     GroupLinkChangeType = "sub_group"
	GroupLinkChangeTypeSibling GroupLinkChangeType = "sibling_group"
)

type GroupUnlinkReason string

const (
	GroupUnlinkReasonDefault GroupUnlinkReason = "unlink_group"
	GroupUnlinkReasonDelete  GroupUnlinkReason = "delete_parent"
)

type GroupLinkTarget struct {
	JID JID
	GroupName
	GroupIsDefaultSub
}

type GroupLinkChange struct {
	Type         GroupLinkChangeType
	UnlinkReason GroupUnlinkReason
	Group        GroupLinkTarget
}
