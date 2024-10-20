# Verification process

## 1. User request

The user provides the DID, the additional security layer password, and the information to be verified (such as age greater than 18 years old).

## 2. Front-end processing

Verify the input format. Encrypt the data and send it to the back-end.

## 3. Back-end verification

The Lambda function verifies the DID and password. Retrieve the user information associated with the DID.

## 4. Zero-knowledge proof generation:

Call the ZKP module to generate a proof based on the stored user information. The proof contains the attributes that need to be verified (such as age), but does not disclose specific information.

## 5. Verification process:

The default verifier or a designated third-party verifier verifies the ZKP.

## 6. Result return:

The backend returns the verification result to the frontend. The frontend displays the verification result to the user.
