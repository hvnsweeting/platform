// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"unicode/utf8"
)

const (
	PREFERENCE_CATEGORY_DIRECT_CHANNEL_SHOW = "direct_channel_show"
	PREFERENCE_CATEGORY_TUTORIAL_STEPS      = "tutorial_step"
	PREFERENCE_CATEGORY_ADVANCED_SETTINGS   = "advanced_settings"
)

type Preference struct {
	UserId   string `json:"user_id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

func (o *Preference) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func PreferenceFromJson(data io.Reader) *Preference {
	decoder := json.NewDecoder(data)
	var o Preference
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func (o *Preference) IsValid() *AppError {
	if len(o.UserId) != 26 {
		return NewAppError("Preference.IsValid", "Invalid user id", "user_id="+o.UserId)
	}

	if len(o.Category) == 0 || len(o.Category) > 32 {
		return NewAppError("Preference.IsValid", "Invalid category", "category="+o.Category)
	}

	if len(o.Name) == 0 || len(o.Name) > 32 {
		return NewAppError("Preference.IsValid", "Invalid name", "name="+o.Name)
	}

	if utf8.RuneCountInString(o.Value) > 128 {
		return NewAppError("Preference.IsValid", "Value is too long", "value="+o.Value)
	}

	return nil
}
