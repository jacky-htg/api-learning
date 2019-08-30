# Rebel Services

Custome Response Status

Make the json standard response contains :
{
    status_code :
    message :
    data :
}

Tasks:
- Change response using standard format [ ResponseFormat ] 
- All call api.Response varying with new response format
- update controllers/tests/userstest.go to varying with new response

## File Changes :
- libraries/api/errors.go
- libraries/api/status_message.go
- libraries/api/status_code.go
- libraries/api/request.go
- libraries/api/response.go
- controllers/users.go
- controllers/tests/userstest.go

## New File :

## Adding Dependency :