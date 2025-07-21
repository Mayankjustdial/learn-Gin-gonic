 Why Modular Structure?
Easier to scale

Promotes separation of concerns

Reusable and testable code

Cleaner project layout for production-grade APIs

📁 Recommended Folder Structure:

/gin-api/
│
├── controllers/        // Business logic handlers
├── routes/             // Route definitions
├── models/             // DB models and structs
├── config/             // Environment & DB setup
├── middlewares/        // Custom middleware
├── utils/              // Helpers (e.g., response utils)
├── main.go             // Entry point
└── go.mod
