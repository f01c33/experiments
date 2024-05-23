-- name: GetUsuario :one
SELECT * FROM usuarios WHERE id = ? LIMIT 1;

-- name: GetUsuarios :many
SELECT * from usuarios ORDER BY nome;

-- name: GetUsuarioPorNome :many
SELECT * from usuarios WHERE nome LIKE ? ORDER BY nome;

-- name: GetUsuarioPorEmail :one
SELECT * from usuarios WHERE email LIKE ? LIMIT 1;

-- name: CreateUsuario :exec
INSERT INTO usuarios (nome,senha,tipo,email) VALUES (?,?,?,?);

-- name: UpdateUsuario :exec
UPDATE usuarios SET nome = ?, email = ?, tipo = ? WHERE id = ?;

-- name: UpdateUsuarioRaw :exec
UPDATE usuarios SET nome = ?, email = ?, tipo = ?, senha = ?, codigo = ? WHERE id = ?;

-- name: DeleteUsuario :exec
DELETE FROM usuarios WHERE id = ?;

-- -- name: GetCasa :one
-- SELECT * FROM casas WHERE id = ? LIMIT 1;

-- -- name: GetCasas :many
-- SELECT * FROM casas ORDER BY nome;

-- -- name: GetCasasPorNome :many
-- SELECT * from casas WHERE nome LIKE ? ORDER BY nome;

-- -- name: GetCasasPorCliente :many
-- SELECT * FROM casas WHERE cliente = ? ORDER BY nome;

-- -- name: CreateCasa :exec
-- INSERT INTO casas (nome,cliente) VALUES (?,?);

-- -- name: UpdateCasa :exec
-- UPDATE casas SET nome=?,cliente=? WHERE id=?;

-- -- name: DeleteCasa :exec
-- DELETE FROM casas WHERE id = ?;

-- -- name: GetCasaCliente :one
-- SELECT * FROM casa_cliente WHERE id = ? LIMIT 1;

-- -- name: GetCasasClientePorCliente :many
-- SELECT * FROM casa_cliente WHERE cliente = ?;

-- -- name: GetCasaClientePorCasa :one
-- SELECT * FROM casa_cliente WHERE casa = ? LIMIT 1;

-- -- name: InsertCasaCliente :exec
-- INSERT INTO casa_cliente (cliente,casa) VALUES (?,?);

-- -- name: UpdateCasaCliente :exec
-- UPDATE casa_cliente SET cliente = ?, casa = ? WHERE id = ?;

-- -- name: DeleteCasaCliente :exec
-- DELETE FROM casa_cliente WHERE id=?;

-- name: GetPost :one
SELECT * FROM posts WHERE id = ? LIMIT 1;

-- name: GetPostsPorCasa :many
SELECT * FROM posts WHERE titulo = ? ORDER BY dt LIMIT ? OFFSET ?;

-- name: GetPosts :many
SELECT * FROM posts ORDER BY dt LIMIT ? OFFSET ?;

-- name: InsertPost :one
INSERT INTO posts (texto,dt,titulo) VALUES (?,?,?) RETURNING *;

-- name: UpdatePost :exec
UPDATE posts SET texto = ?, dt = ?, titulo = ? WHERE id = ?;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = ?;

-- -- name: GetPostPublico :one
-- SELECT * FROM posts_publicos WHERE id = ? LIMIT 1;

-- -- name: GetPostsPublicos :many
-- SELECT * FROM posts_publicos ORDER BY dt LIMIT ? OFFSET ?;

-- -- name: InsertPostPublico :one
-- INSERT INTO posts_publicos (texto,dt) VALUES (?,?) RETURNING *;

-- -- name: UpdatePostPublico :exec
-- UPDATE posts_publicos SET texto = ?, dt = ? WHERE id = ?;

-- -- name: DeletePostPublico :exec
-- DELETE FROM posts_publicos WHERE id = ?;

-- -- name: GetInsumo :one
-- SELECT * FROM insumos WHERE id = ? LIMIT 1;

-- -- name: GetInsumos :many
-- SELECT * FROM insumos ORDER BY texto;

-- -- name: InsertInsumo :exec
-- INSERT INTO insumos (texto,preco) VALUES (?,?);

-- -- name: UpdateInsumo :exec
-- UPDATE insumos SET texto = ?, preco = ? WHERE id = ?;

-- -- name: DeleteInsumo :exec
-- DELETE FROM insumos WHERE id = ?;

-- -- name: GetPostInsumo :one
-- SELECT * FROM post_insumo WHERE id = ? LIMIT 1;

-- -- name: GetPostInsumos :many
-- SELECT * FROM post_insumo ORDER BY id;

-- -- name: GetInsumosForPost :many
-- SELECT insumos.* FROM insumos RIGHT JOIN post_insumo ON post_insumo.post = ? ORDER BY insumos.texto;

-- -- name: InsertPostInsumo :exec
-- INSERT INTO post_insumo (insumo,post) VALUES (?,?);

-- -- name: UpdatePostInsumos :exec
-- UPDATE post_insumo SET insumo = ?, post = ? WHERE id = ?;

-- -- name: DeletePostInsumo :exec
-- DELETE FROM post_insumo WHERE id = ?;

-- -- name: GetRelatorioInsumo :one
-- SELECT * FROM relatorio_insumo WHERE id = ? LIMIT 1;

-- -- name: GetRelatorioInsumos :many
-- SELECT * FROM relatorio_insumo ORDER BY id;

-- -- name: GetInsumosForRelatorio :many
-- SELECT insumos.* FROM insumos RIGHT JOIN relatorio_insumo ON relatorio_insumo.relatorio = ? ORDER BY insumos.texto;

-- -- name: InsertRelatorioInsumo :exec
-- INSERT INTO relatorio_insumo (insumo,relatorio) VALUES (?,?);

-- -- name: UpdateRelatorioInsumos :exec
-- UPDATE relatorio_insumo SET insumo = ?, relatorio = ? WHERE id = ?;

-- -- name: DeleteRelatorioInsumo :exec
-- DELETE FROM relatorio_insumo WHERE id = ?;

-- -- name: GetArquivo :one
-- SELECT * FROM arquivos WHERE id = ? LIMIT 1;

-- -- name: GetArquivos :many
-- SELECT * FROM arquivos ORDER BY dt LIMIT ? OFFSET ?;

-- -- name: GetArquivosCasa :many
-- SELECT * FROM arquivos WHERE casa = ? ORDER BY dt LIMIT ? OFFSET ?;

-- -- name: GetArquivosRelatorio :many
-- SELECT * FROM arquivos WHERE relatorio = ? ORDER BY dt LIMIT ? OFFSET ?;

-- -- name: InsertArquivo :one
-- INSERT INTO arquivos (arquivo, nome, relatorio, dt, casa) VALUES (?,?,?,?,?) RETURNING *;

-- -- name: UpdateArquivo :exec
-- UPDATE arquivos SET arquivo = ?, relatorio = ?, nome = ?, dt = ?, casa = ? WHERE id = ?;

-- -- name: DeleteArquivo :exec
-- DELETE FROM arquivos WHERE id = ?;

-- -- name: CreateRelatorio :one
-- INSERT INTO relatorios (texto,casa,dt) VALUES (?,?,?) RETURNING *;

-- -- name: GetRelatorio :one
-- SELECT * FROM relatorios WHERE id = ? LIMIT 1;

-- -- name: GetRelatorios :many
-- SELECT * FROM relatorios ORDER BY dt LIMIT ? OFFSET ?;

-- -- name: GetRelatoriosCasa :many
-- SELECT * FROM relatorios WHERE casa = ? ORDER BY dt LIMIT ? OFFSET ?;

-- -- name: UpdateRelatorio :exec
-- UPDATE relatorios SET texto = ?, casa = ?, dt = ? WHERE id = ?;

-- -- name: DeleteRelatorio :exec
-- DELETE FROM relatorios WHERE id = ?;

-- name: InsertLogins :exec
INSERT INTO logins (usuario,dt) VALUES (?,?);

-- name: GetLogins :many
SELECT * FROM logins LEFT JOIN usuarios ON logins.usuario=usuarios.id ORDER BY dt LIMIT ? OFFSET ?;