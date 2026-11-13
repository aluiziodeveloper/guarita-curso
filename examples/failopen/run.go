// Package failopen guarda a decisão ingênua do módulo 5: ela nega só
// o que a política nega explicitamente, e repassa todo o resto.
package failopen

import (
	"context"

	"github.com/aluiziodeveloper/guarita-curso/internal/app"
	"github.com/aluiziodeveloper/guarita-curso/internal/domain"
)

// Run falha aberta: decisão desconhecida ou erro da política viram
// repasse ao servidor.
func Run(
	ctx context.Context,
	policy app.PolicyEvaluator,
	upstream app.Upstream,
	call domain.Call,
) (domain.Result, error) {
	decision, _ := policy.Evaluate(ctx, domain.PolicyInput{Call: call})
	if decision == domain.DecisionDeny {
		return domain.Result{}, app.ErrPolicyDenied
	}
	return upstream.Forward(ctx, call)
}
