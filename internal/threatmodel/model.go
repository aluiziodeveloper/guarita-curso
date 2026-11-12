package threatmodel

import (
	"errors"
	"fmt"
	"regexp"
)

// ErrInvalidThreat identifica uma ameaça incompleta no modelo.
var ErrInvalidThreat = errors.New("threat is not valid")

var threatID = regexp.MustCompile(`^[a-z]+(-[a-z]+)*$`)

// Threat é uma linha do modelo de ameaças: o que pode dar errado, por
// onde entra e em que módulo a Guarita ganha a defesa.
type Threat struct {
	ID             string
	Name           string
	Categories     []STRIDECategory
	Boundary       TrustBoundary
	DefenseModule int
	Status         ThreatStatus
}


// Validate exige identificador estável, categorias, fronteira, módulo
// e estado válidos.
func (threat Threat) Validate() error {
	if !threatID.MatchString(threat.ID) || threat.Name == "" {
		return fmt.Errorf("threat %q: %w", threat.ID, ErrInvalidThreat)
	}
	if len(threat.Categories) == 0 {
		return fmt.Errorf("threat %q: no STRIDE category: %w",
			threat.ID, ErrInvalidThreat)
	}
	for _, category := range threat.Categories {
		if err := category.Validate(); err != nil {
			return fmt.Errorf("threat %q: %w", threat.ID, err)
		}
	}
	if err := threat.Boundary.Validate(); err != nil {
		return fmt.Errorf("threat %q: %w", threat.ID, err)
	}
	if threat.DefenseModule < 1 || threat.DefenseModule > 23 {
		return fmt.Errorf("threat %q: module %d: %w",
			threat.ID, threat.DefenseModule, ErrInvalidThreat)
	}
	if err := threat.Status.Validate(); err != nil {
		return fmt.Errorf("threat %q: %w", threat.ID, err)
	}
	return nil
}

// Model devolve uma cópia nova do modelo a cada chamada, para que
// nenhum chamador altere o modelo visto por outro.
func Model() []Threat {
	in, out := TrustZoneMCPServer, TrustZoneGuarita
	agent, operator := TrustZoneAgent, TrustZoneOperator
	return []Threat{
		threat("tool-poisoning", "Envenenamento de ferramenta",
			edge(in, out), 11, tampering, elevationOfPrivilege),
		threat("rug-pull", "Troca silenciosa da definição aprovada",
			edge(in, out), 10, tampering),
		threat("tool-shadowing",
			"Sombreamento de ferramenta por nome repetido",
			edge(in, out), 11, spoofing),
		threat("indirect-prompt-injection",
			"Injeção de prompt indireta na resposta",
			edge(in, out), 15, tampering, elevationOfPrivilege),
		threat("excessive-agency", "Agência excessiva",
			edge(agent, out), 12, elevationOfPrivilege),
		threat("data-exfiltration",
			"Exfiltração de dados nos argumentos",
			edge(out, in), 14, informationDisclosure),
		threat("sandbox-escape", "Fuga da sandbox do servidor",
			edge(out, in), 18, elevationOfPrivilege),
		threat("credential-theft",
			"Repasse ou roubo de credenciais (confused deputy)",
			edge(out, in), 9, spoofing, informationDisclosure),
		threat("supply-chain", "Pacote de servidor comprometido",
			edge(operator, in), 21, tampering),
		threat("resource-abuse",
			"Laços de chamadas que esgotam cota ou custo",
			edge(agent, out), 19, denialOfService),
		threat("insecure-parsing",
			"Mensagem malformada ou ambígua no protocolo",
			edge(agent, out), 7, tampering, denialOfService),
		threat("approval-forgery",
			"Aprovação forjada ou obtida por fadiga",
			edge(operator, out), 13, spoofing, repudiation),
	}
}


// Nomes curtos para as categorias, só para manter o modelo legível.
const (
	spoofing              = STRIDESpoofing
	tampering             = STRIDETampering
	repudiation           = STRIDERepudiation
	informationDisclosure = STRIDEInformationDisclosure
	denialOfService       = STRIDEDenialOfService
	elevationOfPrivilege  = STRIDEElevationOfPrivilege
)

func threat(
	id, name string,
	boundary TrustBoundary,
	module int,
	categories ...STRIDECategory,
) Threat {
	return Threat{
		ID:             id,
		Name:           name,
		Categories:     categories,
		Boundary:       boundary,
		DefenseModule: module,
		Status:         ThreatStatusOpen,
	}
}

func edge(from, to TrustZone) TrustBoundary {
	return TrustBoundary{From: from, To: to}
}
