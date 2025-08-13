# Melhorias Futuras - API Alemão

Este arquivo contém ideias e planos para evoluir a API atual, tornando-a mais completa, segura e preparada para produção.

---

## 🔒 Autenticação e Segurança

- [ ] Criar sistema completo de usuários com cadastro (`/signup`)
- [ ] Validar se o e-mail já existe antes de cadastrar
- [ ] Adicionar refresh token (opcional, mais avançado)
- [ ] Armazenar `hash` das senhas com `bcrypt`
- [ ] Impedir que um usuário acesse frases de outro (segurança por ID)

---

## ⚙️ Funcionalidades

- [ ] Adicionar rota `GET /phrases/:id` para buscar frase específica
- [ ] Adicionar paginação em `GET /phrases`
- [ ] Adicionar campos como `created_at`, `updated_at` (timestamps)
- [ ] Permitir que o usuário "curta" ou "favorite" uma frase
- [ ] Criar uma rota pública que mostra frases populares aleatórias

---

## 📦 Organização e Boas Práticas

- [ ] Separar o projeto em camadas (`handlers`, `services`, `models`, `routes`, `middleware`)
- [ ] Usar variáveis de ambiente reais com `.env` (ex: `godotenv`)
- [ ] Adicionar logs mais informativos
- [ ] Criar middleware global para tratamento de erros

---

## 📄 Documentação

- [ ] Integrar Swagger para documentação automática da API
- [ ] Escrever README mais completo com exemplos de uso (curl/Postman)
- [ ] Documentar o modelo de dados e regras de validação

---

## 🧪 Testes

- [ ] Adicionar testes unitários para os handlers e serviços
- [ ] Adicionar testes de integração (ex: usando banco em memória)
- [ ] Automatizar testes antes de cada push/deploy

---

## 🚀 Deploy

- [ ] Criar Dockerfile para rodar a API em container
- [ ] Configurar docker-compose com banco de dados
- [ ] Fazer deploy em ambiente gratuito (Render, Railway, etc.)
- [ ] Adicionar monitoramento básico (logs, uptime)

---

Esses itens podem ser feitos aos poucos, de forma modular. Você pode ir marcando os que já foram feitos ✅.
