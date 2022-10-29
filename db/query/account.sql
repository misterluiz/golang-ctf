-- name: CreateAccount :one 
INSERT INTO accounts(
    user_id,
    category_id,
    tytle,
    type,
    description,
    value,
    date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7  
)RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccounts :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM 
    accounts a
LEFT JOIN 
    categories c ON c.id = a.category_id 
WHERE 
    a.user_id = @user_id
AND
    a.type = @type
AND
    LOWER(a.tytle) LIKE CONCAT('%', LOWER(@tytle::text), '%')
AND 
    LOWER(a.description) LIKE CONCAT('%', LOWER(@description::text), '%')
AND
    a.category_id = COALESCE(sqlc.narg('category_id'), a.category_id)
AND 
    a.date = COALESCE(sqlc.narg('date'), a.date);





-- name: GetAccountsByUserIdAndType :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2;

-- name: GetAccountsByUserIdAndTypeAndCategoryId :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3 ;


-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTytle :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3 AND a.tytle LIKE $4;


-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTytleAndDescription :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 
AND a.category_id = $3 AND a.tytle LIKE $4 AND a.description LIKE $5;

-- name: GetAccountsReports :one 
SELECT SUM(value) AS sum_value FROM accounts
WHERE user_id = $1 AND type = $2;

-- name: GetAccountsGraph :one 
SELECT COUNT(*) FROM accounts
WHERE user_id = $1 AND type = $2;

-- name: UpdateAccount :one
UPDATE accounts SET tytle = $2, description = $3, 
value = $4 WHERE id = $1 RETURNING *; 

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE  id = $1;



-- name: GetAccountsByUserIdAndTypeAndTytle :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 AND a.tytle LIKE $3;

-- name: GetAccountsByUserIdAndTypeAndDescription :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 AND a.description LIKE $3;

-- name: GetAccountsByUserIdAndTypeAndDate :many 
SELECT 
 a.id,
 a.user_id,
 a.tytle,
 a.type,
 a.description,
 a.value,
 a.date,
 a.created_at,
 c.tytle as category_tytle
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 AND a.date LIKE $3;