package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/liveleaderboard"

	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/database"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/env"
	"github.com/Excel-MEC/excelplay-backend-kryptos/pkg/httperrors"
	"github.com/dgrijalva/jwt-go"
)

// Only for swagger documentation, do not use in code.
type swagRequest struct {
	Answer string `json:"answer" example:"excel"`
}

// HandleSubmission handles answer attempts
// @Summary takes a post request with the answer attempt.
// @Description takes a post request with the answer attempt.
// @Tags Kryptos
// @Accept json
// @Produce plain
// @Param payload body swagRequest true "Answer format"
// @Success 200 {object} string "Returns 'success' for correct answer, 'fail' for wrong answer."
// @Failure 500 {string} string
// @Router /api/submit [post]
func HandleSubmission(db *database.DB, env *env.Config) httperrors.Handler {
	type request struct {
		Answer string `json:"answer"`
	}
	return func(w http.ResponseWriter, r *http.Request) *httperrors.HTTPError {
		// Obtain values from JWT
		props, _ := r.Context().Value("props").(jwt.MapClaims)
		userID, _ := strconv.Atoi(props["user_id"].(string))

		// Expected POST format is { "answer": "attempt" }
		input := json.NewDecoder(r.Body)
		input.DisallowUnknownFields()

		var req request
		err := input.Decode(&req)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not deserialize json", http.StatusInternalServerError}
		}

		var currUser database.User
		err = db.GetUser(&currUser, userID)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve user", http.StatusInternalServerError}
		}

		_, err = db.LogAnswerAttempt(userID, currUser, req.Answer)

		var correctAns string
		err = db.GetCorrectAns(currUser, &correctAns)
		if err != nil {
			return &httperrors.HTTPError{r, err, "Could not retrieve the answer", http.StatusInternalServerError}
		}

		if req.Answer == correctAns {
			_, err := db.CorrectAnswerSubmitted(userID)
			if err != nil {
				return &httperrors.HTTPError{r, err, "Could not update user progress", http.StatusInternalServerError}
			}
			// Send update to leaderboard
			liveleaderboard.UpdateUser <- userID
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
			return nil
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("fail"))
		return nil
	}
}
