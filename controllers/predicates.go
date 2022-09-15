/*
Copyright 2020 The Flux authors

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

package controllers

import (
	kluctlv1 "github.com/kluctl/flux-kluctl-controller/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"
)

type SourceRevisionChangePredicate struct {
	predicate.Funcs
}

func (SourceRevisionChangePredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld == nil || e.ObjectNew == nil {
		return false
	}

	oldSource, ok := e.ObjectOld.(sourcev1.Source)
	if !ok {
		return false
	}

	newSource, ok := e.ObjectNew.(sourcev1.Source)
	if !ok {
		return false
	}

	if oldSource.GetArtifact() == nil && newSource.GetArtifact() != nil {
		return true
	}

	if oldSource.GetArtifact() != nil && newSource.GetArtifact() != nil &&
		oldSource.GetArtifact().Revision != newSource.GetArtifact().Revision {
		return true
	}

	return false
}

type DeployRequestedPredicate struct {
	predicate.Funcs
}

func (DeployRequestedPredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld == nil || e.ObjectNew == nil {
		return false
	}

	if val, ok := e.ObjectNew.GetAnnotations()[kluctlv1.KluctlDeployRequestAnnotation]; ok {
		if valOld, okOld := e.ObjectOld.GetAnnotations()[kluctlv1.KluctlDeployRequestAnnotation]; okOld {
			return val != valOld
		}
		return true
	}
	return false
}
