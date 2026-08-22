# Quantum Circuit Playground & Template Engine — Backend

Serviço web de alta performance desenvolvido em Go para criação, gerenciamento e simulação interativa de circuitos quânticos personalizados. O projeto foi desenhado do zero aplicando princípios rígidos de **Clean Architecture**, **Clean Code** e **Secure by Design**.

---

## 🎯 Sobre o Projeto

Inspirado em plataformas visuais e interativas como o *TensorFlow Playground*, este backend é responsável por:
1. **Engine de Simulação Quântica:** Processamento de álgebra linear e multiplicação de portões quânticos (Hadamard, CNOT, Pauli-X/Y/Z) sobre vetores de estado de qubits.
2. **Gerador de Templates de Circuitos:** Permite que usuários criem, versionem e compartilhem estruturas reutilizáveis de circuitos quânticos.
3. **API RESTful de Baixa Latência:** Exposição segura dos serviços para o frontend interativo.

---

## 🚀 Stack Tecnológica

| Camada | Tecnologia | Justificativa |
| :--- | :--- | :--- |
| **Linguagem** | Go (1.22+) | Alta performance compilada, concorrência nativa e baixo overhead de memória para simulações. |
| **HTTP / Router** | `net/http` + `chi` | Roteamento leve, idiomático e com suporte nativo a middlewares de segurança. |
| **Banco de Dados** | PostgreSQL | Persistência confiável de dados relacionais (usuários, templates e históricos). |
| **Acesso a Dados** | `sqlc` + `pgx` | Geração de código Go *type-safe* a partir de SQL puro. Zero ORM, zero reflection e máxima performance. |
| **Autenticação** | JWT (`golang-jwt`) | Autenticação stateless, segura e com controle estrito de claims. |
| **Validação** | `go-playground/validator` | Validação declarativa e fortemente tipada das entradas de API. |
| **Logs** | `slog` (Stdlib) | Log estruturado em JSON para rastreabilidade de requisições e auditoria. |
| **Documentação** | Swagger (`swag`) | Documentação automática e interativa das APIs gerada via anotações no código. |
| **Testes** | `testing` + `testify` | Suíte de testes unitários, de integração e de benchmarking nativos. |

---

## 🛡️ Arquitetura e Boas Práticas

### Clean Architecture
A aplicação é dividida em camadas bem definidas, garantindo o desacoplamento do domínio quântico de detalhes de infraestrutura:

```text
cmd/api/               # Ponto de entrada da aplicação
internal/
├── domain/            # Entidades puras (Qubit, Portões Quânticos, Circuito, User)
├── usecase/           # Regras de negócio e motores de simulação
├── handler/           # Adaptadores HTTP / REST Controllers
└── repository/        # Camada de persistência (sqlc / in-memory / postgres)