-- name: CreateWallet :one
INSERT INTO wallets (
  private_key, public_key
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets 
WHERE public_key = $1
LIMIT 1;