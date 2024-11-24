# Quake Log Parser

## 📋 Motivação

Este projeto foi desenvolvido para processar o arquivo de log do Quake (`qgames.log`) e extrair informações relevantes para gerar relatórios personalizados. Ele permite analisar dados como agrupamento de partidas e estatísticas de mortes.

---

## 🔧 Descrição Técnica

O projeto foi implementado em **Go (Golang)** e oferece as seguintes funcionalidades:

- **📂 Leitura do Arquivo de Log**  
  Processa o log e armazena o resultado em `reading_the_log_file.json`.

- **🗂️ Agrupamento de Dados por Partida**  
  Organiza os dados por jogo e gera o arquivo `grouping_data_by_game.json`.

- **⚔️ Coleta de Dados de Mortes**  
  Analisa as mortes no jogo e registra os resultados em `deaths_by_means_game.json`.

### ⚙️ Automação com Makefile

### Pré-requisitos

Antes de utilizar o **Makefile**, verifique se o **make** está instalado em sua máquina. Caso contrário, você pode instalá-lo utilizando o gerenciador de pacotes da sua plataforma. Por exemplo:

Um **Makefile** foi criado para facilitar o uso de comandos recorrentes, como:

- **🧪 Testes de Unidade e Relatórios**  
  Execute todos os testes e gere relatórios automaticamente. Para executar os testes basta digitar o comando:
```bash
  make test-report
```

- **📦 Build Multi-OS**  
  Gere executáveis para os principais sistemas operacionais:
  - Linux
  - Windows
  - macOS

Para criar os artefatos executáveis, basta usar o comando:

```bash
make build
```

### 🚀 Pipeline de Integração Contínua (CI)

Para garantir um fluxo de trabalho ágil e eficiente, o projeto utiliza um **Pipeline de Integração Contínua (CI)**. Esse pipeline automatiza etapas cruciais do desenvolvimento, como:

- **Execução de Testes Automatizados**  
  Todos os testes de unidade são executados a cada commit para garantir que novas alterações não quebrem funcionalidades existentes.

- **Build Automatizado**  
  Realiza o build do projeto para todas as plataformas suportadas (**Linux**, **Windows** e **macOS**), validando que o código está funcional em diferentes ambientes.

- **Notificação de Falhas**  
  Em caso de falhas nos testes ou na etapa de build, o pipeline notifica a equipe, permitindo uma resposta rápida para corrigir os problemas.

Esse processo automatizado promove **agilidade** e **confiança** no desenvolvimento, garantindo que o código entregue seja de alta qualidade.