# Deploy RDS using Serverless
echo "Deploying RDS using Serverless..."
serverless deploy --config serverless-rds.yml

# Deploy Lambda function using Serverless
echo "Deploying Lambda function using Serverless..."
serverless deploy --config serverless-lambda.yml

# Retrieve database URL from CloudFormation stack outputs
echo "Retrieving database URL..."
DB_URL=$(aws cloudformation describe-stacks --stack-name rds-setup-4-dev --query "Stacks[0].Outputs[?OutputKey=='DatabaseURL'].OutputValue" --output text)
export DATABASE_URL=${DB_URL}

# Run database migrations using Goose
echo "Running database migrations..."
goose -dir src/infrastructure/database/migrations postgres "$DATABASE_URL" up

echo "Deployment completed successfully!"