package router

import (
	"net/http"

	"github.com/openmeterio/openmeter/api"
)

// List plans
// (GET /api/v1/plans)
func (a *Router) ListPlans(w http.ResponseWriter, r *http.Request, params api.ListPlansParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a plan
// (POST /api/v1/plans)
func (a *Router) CreatePlan(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete plan
// (DELETE /api/v1/plans/{planId})
func (a *Router) DeletePlan(w http.ResponseWriter, r *http.Request, planId string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get plan
// (GET /api/v1/plans/{planId})
func (a *Router) GetPlan(w http.ResponseWriter, r *http.Request, planId string, params api.GetPlanParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a plan
// (PUT /api/v1/plans/{planId})
func (a *Router) UpdatePlan(w http.ResponseWriter, r *http.Request, planId string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// New draft plan
// (POST /api/v1/plans/{planId}/next)
func (a *Router) NextPlan(w http.ResponseWriter, r *http.Request, planId string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List phases in plan
// (GET /api/v1/plans/{planId}/phases)
func (a *Router) ListPlanPhases(w http.ResponseWriter, r *http.Request, planId string, params api.ListPlanPhasesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create new phase in plan
// (POST /api/v1/plans/{planId}/phases)
func (a *Router) CreatePlanPhase(w http.ResponseWriter, r *http.Request, planId string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete phase for plan
// (DELETE /api/v1/plans/{planId}/phases/{planPhaseKey})
func (a *Router) DeletePlanPhase(w http.ResponseWriter, r *http.Request, planId string, planPhaseKey string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get phase for plan
// (GET /api/v1/plans/{planId}/phases/{planPhaseKey})
func (a *Router) GetPlanPhase(w http.ResponseWriter, r *http.Request, planId string, planPhaseKey string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update phase in plan
// (PUT /api/v1/plans/{planId}/phases/{planPhaseKey})
func (a *Router) UpdatePlanPhase(w http.ResponseWriter, r *http.Request, planId string, planPhaseKey string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Publish plan
// (POST /api/v1/plans/{planId}/publish)
func (a *Router) PublishPlan(w http.ResponseWriter, r *http.Request, planId string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Archive plan version
// (POST /api/v1/plans/{planId}/archive)
func (a *Router) ArchivePlan(w http.ResponseWriter, r *http.Request, planId string) {
	w.WriteHeader(http.StatusNotImplemented)
}
