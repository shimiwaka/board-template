# board-template

template for simple web service

## how to use

If you create a service called SERVICE_NAME, exec it.

`find . -type f | xargs sed -i 's/board-template/SERVICE_NAME/g'`

## test command

`curl -X POST --data-urlencode 'email=hoge@email.com' http://localhost:9999/create`
