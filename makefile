remockgen:
	@echo "===Remockgen==="
	@echo "==REPOSITORY=="
	@mockgen -source=./repository/database/otp.go -destination=./mock/repository/database/mock_otp.go
	@mockgen -source=./repository/database/user.go -destination=./mock/repository/database/mock_user.go
	@mockgen -source=./usecase/otp.go -destination=./mock/usecase/mock_otp.go
	@mockgen -source=./usecase/user.go -destination=./mock/usecase/mock_user.go