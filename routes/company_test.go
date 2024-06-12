package routes

import "testing"

func Test_PostCompanyRequestToCompany_TestSlugs(t *testing.T) {
	var tests = []struct {
		Name          string
		PCR           PostCompanyRequest
		ExpectedError error
	}{
		{
			PCR: PostCompanyRequest{
				Name: "",
			},
			ExpectedError: ErrorInvalidSlugLength,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := PostCompanyRequestToCompany(tt.PCR)

			// can ignore wrapped error warning because we don't used wrapped errors atm
			if err != tt.ExpectedError {
				t.Errorf("got %v expected %v", err, tt.ExpectedError)
			}
		})
	}
}
