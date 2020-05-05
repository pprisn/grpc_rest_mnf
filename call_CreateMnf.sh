#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service MnfService cli --call CreateMnf --file RequestCreateMnf.json ClientStreaming
#OUTPUT NEW ID
#{
#  "id": "7da4bd51-67b6-4b82-ad48-fa2bfa37a503"
#}

