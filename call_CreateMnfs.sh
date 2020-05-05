#!/bin/sh
evans  --proto api/mnf/v1/mnf_.proto --port 2338 --service MnfService cli --call CreateMnfs --file RequestCreateMnfs.json ClientStreaming
#OUTPUT NEW []ID
{
  "ids": [
    "971a3522-47f3-4dd6-b075-5c60f3f60136",
    "35ed13f5-aca4-4b33-b9e1-aad4023d6d77",
    "ca04e4de-654d-4814-8ba8-6b53da7b7724",
    "390a13d0-ca93-4766-809b-a9994b964f38"
  ]
}

