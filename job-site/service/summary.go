package service

import "job-post/models"

func (Service) Summary(jobResp []models.PostResponse, summary *models.Summary) error {
	summary.TotalJobs = len(jobResp)

	jobs := make(map[string]int)
	countries := make(map[string]int)

	for _, value := range jobResp {
		jobs[value.JobTitle] = jobs[value.JobTitle] + 1
		countries[value.Country.Country] = countries[value.Country.Country] + 1
	}

	for key := range jobs {
		job := models.Jobs{
			Name:  key,
			Total: jobs[key],
		}
		summary.Jobs = append(summary.Jobs, job)
	}

	summary.Countries = len(countries)
	return nil
}
