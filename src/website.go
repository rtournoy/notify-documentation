package main

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Website struct {
	WebrootFolderPath       string
	ContentFolderPath       string
	ScenariosFolderPath     string
	PatternsFolderPath      string
	StaticContentFolderPath string
	Scenarios               []*Scenario
	Patterns                []*Pattern
	//NotificationsFolderPath        string
}

func (website *Website) Initialise(webrootPath, scenariosCsvSource, workflowStepsCsvSource, patternsCsvSource string) error {
	var err error
	website.WebrootFolderPath = webrootPath
	website.ContentFolderPath = filepath.Join(website.WebrootFolderPath, "content")
	website.ScenariosFolderPath = filepath.Join(website.ContentFolderPath, "scenarios")
	website.PatternsFolderPath = filepath.Join(website.ContentFolderPath, "patterns")
	//website.NotificationsFolderPath = filepath.Join(website.ContentFolderPath, "notifications")
	website.StaticContentFolderPath = filepath.Join(website.WebrootFolderPath, "static")
	scenariosCsvInput, err := readBytesFromFileOrUrl(scenariosCsvSource)
	if err != nil {
		return err
	}
	err = csvutil.Unmarshal(scenariosCsvInput, &website.Scenarios)
	if err != nil {
		return err
	}
	zapLogger.Info("Scenarios loaded OK")
	patternsCsvInput, err := readBytesFromFileOrUrl(patternsCsvSource)
	if err != nil {
		return err
	}
	err = csvutil.Unmarshal(patternsCsvInput, &website.Patterns)
	if err != nil {
		return err
	}
	zapLogger.Info("Patterns loaded OK")
	workflowStepsCsvInput, err := readBytesFromFileOrUrl(workflowStepsCsvSource)
	if err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	var workflowSteps []WorkflowStep
	err = csvutil.Unmarshal(workflowStepsCsvInput, &workflowSteps)
	if err != nil {
		zapLogger.Error(err.Error())
		return err
	}
	for _, workflowStep := range workflowSteps {
		scenario := website.GetScenarioById(workflowStep.ScenarioId)
		if scenario != nil {
			scenario.Notifications = append(scenario.Notifications, workflowStep)
		}
	}
	return err
}

func (website *Website) GetScenarioById(id string) *Scenario {
	for _, scenario := range website.Scenarios {
		if scenario.Id == id {
			return scenario
		}
	}
	return nil
}

func (website *Website) WriteWebPages() error {
	var err error
	for _, scenario := range website.Scenarios {
		zapLogger.Debug("Scenario:", zap.String("ID", scenario.Id))
		scenarioFolderPath := filepath.Join(website.ScenariosFolderPath, scenario.Id)
		zapLogger.Debug("Scenario folder path:", zap.String("Path", scenarioFolderPath))
		err = resetVolatileFolder(scenarioFolderPath)
		if err != nil {
			zapLogger.Error(err.Error())
			return err
		}
		scenarioPage, scenarioPageMarshallErr := scenario.Marshal()
		if scenarioPageMarshallErr != nil {
			zapLogger.Error(scenarioPageMarshallErr.Error())
			return scenarioPageMarshallErr
		}
		scenarioPageWriteErr := ioutil.WriteFile(filepath.Join(scenarioFolderPath, "_index.md"), scenarioPage, os.ModePerm)
		if scenarioPageWriteErr != nil {
			zapLogger.Error(scenarioPageWriteErr.Error())
			return scenarioPageWriteErr
		}
		for _, notification := range scenario.Notifications {
			notificationFilePath := filepath.Join(scenarioFolderPath, fmt.Sprintf("%v.md", notification.Position))
			notificationPage, notificationMarshallErr := notification.Marshal()
			if notificationMarshallErr != nil {
				zapLogger.Error(notificationMarshallErr.Error())
				return notificationMarshallErr
			}
			notificationPageWriteErr := ioutil.WriteFile(notificationFilePath, notificationPage, os.ModePerm)
			if notificationPageWriteErr != nil {
				zapLogger.Error(notificationPageWriteErr.Error())
				return notificationPageWriteErr
			}
		}
	}
	for _, pattern := range website.Patterns {
		zapLogger.Debug("Pattern:", zap.String("ID", pattern.Id))
		patternFolderPath := filepath.Join(website.PatternsFolderPath, pattern.Id)
		zapLogger.Debug("Pattern folder path:", zap.String("Path", patternFolderPath))
		err = resetVolatileFolder(patternFolderPath)
		if err != nil {
			zapLogger.Error(err.Error())
			return err
		}
		patternPage, patternPageMarshallErr := pattern.Marshal()
		if patternPageMarshallErr != nil {
			zapLogger.Error(patternPageMarshallErr.Error())
			return patternPageMarshallErr
		}
		patternPageWriteErr := ioutil.WriteFile(filepath.Join(patternFolderPath, "index.md"), patternPage, os.ModePerm)
		if patternPageWriteErr != nil {
			zapLogger.Error(patternPageWriteErr.Error())
			return patternPageWriteErr
		}
	}
	return err
}
