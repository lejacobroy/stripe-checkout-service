# Stripe Checkout Service

This simple service includes one endpoint, which accepts the outputted Stripe checkout token and charges it.

This repo is ready to deploy to Heroku, you just need to set a few environment variables:


Required | ENV | example
---------|-----|--------
yes | `SECRET_KEY` | sk_test_UvLkbnT99KVPEMb7ua1M5afN
yes | `CHARGE_DESCRIPTION` | Personal gift to Jacob
no  | `CURRENCY` | usd


Then You can post a few parameters to `/charge` as `Content-Type: application/x-www-form-urlencoded` data.

Key | Value description
----|------------------
`email` | The customer email collected in checkout
`id` | The token identifier returned by checkout.js
`amount` | How much of the charge to capture
