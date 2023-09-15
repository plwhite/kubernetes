/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package podnetworkattachment

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/networking"
	"k8s.io/kubernetes/pkg/apis/networking/validation"
)

// podNetworkAttachmentStrategy implements verification logic for PodNetworkAttachments.
type podNetworkAttachmentStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// Strategy is the default logic that applies when creating and updating PodNetworkAttachment objects.
var Strategy = podNetworkAttachmentStrategy{legacyscheme.Scheme, names.SimpleNameGenerator}

// NamespaceScoped returns false because all podNetworkAttachments do not need to be within a namespace.
func (podNetworkAttachmentStrategy) NamespaceScoped() bool {
	return false
}

func (podNetworkAttachmentStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {}

func (podNetworkAttachmentStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {}

// Validate validates a new PodNetworkAttachments.
func (podNetworkAttachmentStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	podNetworkAttachment := obj.(*networking.PodNetworkAttachment)
	return validation.ValidatePodNetworkAttachment(podNetworkAttachment)
}

// WarningsOnCreate returns warnings for the creation of the given object.
func (podNetworkAttachmentStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

// Canonicalize normalizes the object after validation.
func (podNetworkAttachmentStrategy) Canonicalize(obj runtime.Object) {}

// AllowCreateOnUpdate is false for PodNetworkAttachment; this means POST is needed to create one.
func (podNetworkAttachmentStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (podNetworkAttachmentStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	validationErrorList := validation.ValidatePodNetworkAttachment(obj.(*networking.PodNetworkAttachment))
	updateErrorList := validation.ValidatePodNetworkAttachmentUpdate(obj.(*networking.PodNetworkAttachment), old.(*networking.PodNetworkAttachment))
	return append(validationErrorList, updateErrorList...)
}

// WarningsOnUpdate returns warnings for the given update.
func (podNetworkAttachmentStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

// AllowUnconditionalUpdate is the default update policy for PodNetworkAttachment objects.
func (podNetworkAttachmentStrategy) AllowUnconditionalUpdate() bool {
	return true
}
