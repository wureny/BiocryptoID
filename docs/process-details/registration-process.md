# Registration process

## 1. User input

Users upload facial images through the front-end interface.

Enter a custom additional security layer password.

Provide personal information: name, date of birth, gender, etc.

> Facial information is required for more than 1 week

## 2. Front-end processing:

React front-end performs preliminary verification of input.

Package and encrypt data and send it to the back-end through a secure channel.

## 3. Back-end processing:

AWS API Gateway receives the request. The Lambda function is triggered and starts processing the request.

## 4. Identity creation:

Generate a unique distributed identity identifier (DID). Associate the DID with the user information.

## 5. Biometric processing:

Call the AWS Recognition API to process the face image. Extract and encrypt the facial feature data.

## 6. Data storage:

Store the encrypted facial feature data in AWS S3 with the DID as the key. Other user information is stored in a secure way.

## 7. Confirmation:

Return a successful registration message and DID to the user.
