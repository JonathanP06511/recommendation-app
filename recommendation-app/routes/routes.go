package routes

import (
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
    "recommendation-app/models"
)

func RecommendationsHandler(w http.ResponseWriter, r *http.Request) {
    recommendations := []models.Recommendation{
        {ID: 1, Text: "Drink enough water every day."},
        {ID: 2, Text: "Sleep at least 8 hours each night."},
        {ID: 3, Text: "Exercise regularly to maintain a healthy body."},
        {ID: 4, Text: "Eat a balanced diet rich in fruits and vegetables."},
        {ID: 5, Text: "Practice mindfulness to reduce stress."},
        {ID: 6, Text: "Get regular medical check-ups."},
        {ID: 7, Text: "Limit your intake of sugary drinks and snacks."},
        {ID: 8, Text: "Wash your hands frequently to prevent infections."},
        {ID: 9, Text: "Avoid smoking and limit alcohol consumption."},
        {ID: 10, Text: "Take breaks and stretch during long periods of sitting."},
        {ID: 11, Text: "Incorporate physical activities you enjoy into your routine."},
        {ID: 12, Text: "Ensure proper posture to avoid back and neck pain."},
        {ID: 13, Text: "Spend time outdoors to get fresh air and sunlight."},
        {ID: 14, Text: "Avoid processed foods; choose whole foods instead."},
        {ID: 15, Text: "Stay socially connected to maintain emotional health."},
        {ID: 16, Text: "Eat slowly and mindfully to aid digestion."},
        {ID: 17, Text: "Get enough sleep by maintaining a consistent sleep schedule."},
        {ID: 18, Text: "Practice deep breathing exercises to improve lung function."},
        {ID: 19, Text: "Keep a positive attitude and practice gratitude daily."},
        {ID: 20, Text: "Include sources of healthy fats in your diet, such as nuts and avocados."},
    }

    rand.Seed(time.Now().UnixNano())
    randomIndex := rand.Intn(len(recommendations))
    recommendation := recommendations[randomIndex]

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(recommendation)
}
