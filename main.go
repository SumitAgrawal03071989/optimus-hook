package main

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/plugin"

)

var (
	Name    = "smiley"
	Version = "latest"
	Image   = "dockerhub.com/sumitagrawal03071989/smiley:0.1"
)

type Smiley struct{}



func (p *Smiley) PluginInfo() (*models.PluginInfoResponse, error) {
	return &models.PluginInfoResponse{
		Name:          "SmileyPrinter",
		Description:   "JustPrintsTheSmileyFace",
		//Image:         fmt.Sprintf("%s:%s", Image, Version),
		DependsOn:     []string{},
		PluginType:    models.PluginTypeHook,
		//PluginVersion: Version,
		HookType:      models.HookTypePre,
		PluginMods:    []models.PluginMod{models.ModTypeCLI},
	}, nil
}

func (p *Smiley) GetQuestions(ctx context.Context, req models.GetQuestionsRequest) (*models.GetQuestionsResponse, error) {
	return &models.GetQuestionsResponse{
		Questions: nil,
	}, nil
}


func (t *Smiley) DefaultAssets(ctx context.Context, _ models.DefaultAssetsRequest) (*models.DefaultAssetsResponse, error) {
	return &models.DefaultAssetsResponse{}, nil
}

func (t *Smiley) CompileAssets(ctx context.Context, req models.CompileAssetsRequest) (*models.CompileAssetsResponse, error) {
	return &models.CompileAssetsResponse{
		Assets: req.Assets,
	}, nil
}


func (p *Smiley) ValidateQuestion(ctx context.Context, req models.ValidateQuestionRequest) (*models.ValidateQuestionResponse, error) {
	return &models.ValidateQuestionResponse{
		Success: true,
	}, nil
}


func (p *Smiley) DefaultConfig(ctx context.Context, request models.DefaultConfigRequest) (*models.DefaultConfigResponse, error) {
	conf := []models.PluginConfig{
		{
			Name:  "Type of smiley faces you want to print",
			Value: `Any`,
		},
	}
	return &models.DefaultConfigResponse{
		Config: conf,
	}, nil
}

func main() {
	plugin.Serve(func(log hclog.Logger) interface{} {
		return &Smiley{
		}
	})
}
