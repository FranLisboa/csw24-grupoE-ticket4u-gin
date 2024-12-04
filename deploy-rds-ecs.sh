echo "Deploying RDS using Serverless..."
serverless deploy --config serverless-rds.yml

echo "Retrieving database URL..."
DB_URL=$(aws cloudformation describe-stacks --stack-name rds-setup-4-dev --query "Stacks[0].Outputs[?OutputKey=='DatabaseURL'].OutputValue" --output text)
export DATABASE_URL=${DB_URL}

echo "Running database migrations..."
goose -dir src/infrastructure/database/migrations postgres "$DATABASE_URL" up

echo "initializing rcs terraform"
cd infra

terraform init

terraform apply -var="DATABASE_URL=$DATABASE_URL" -var="AWS_ROLE=$AWS_ROLE"J