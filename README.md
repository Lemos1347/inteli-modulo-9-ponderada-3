# Simulador de Dispositivos IoT integrado a um broker remoto

## 1. Objetivo

Este projeto visa criar um simulador de dispositivos IoT que utiliza o protocolo MQTT para enviar informações simuladas baseadas em dados de sensores reais, nesse caso o [Sensor de Radiação Solar Sem Fio HOBOnet RXW-LIB-900](https://sigmasensors.com.br/produtos/sensor-de-radiacao-solar-sem-fio-hobonet-rxw-lib-900).

## 2. Como Instalar e Rodar

### Pré-requisitos

- Credenciais de um broker remoto que tenha conexão TLS (recomendamos [HiveMQ](https://www.hivemq.com/))

> [!IMPORTANT]
> Caso queira editar o código ou rodá-lo de outra maneira, tenha instalado [go](https://go.dev/doc/install) em sua máquina.

### Instalação

Clone o repositório para a sua máquina local:

```bash
git clone https://github.com/Lemos1347/inteli-modulo-9-ponderada-3
cd inteli-modulo-9-ponderada-3
```

> [!NOTE]
> Caso queira editar o projeto, lembre de instalar as depencias de cada módulo:
>
> ```bash
> cd publisher ; go mod tidy ; cd ../subscriber ; go mod tidy ; cd ..
> ```

### Execução

Crie um arquivo `.env` na raiz do projeto. Esse arquivo deve ter o seguinte formato e você deve preencher com as credenciais do seu broker:

```env
BROKER_URL=<URL DO SEU BROKER>
MQTT_PUB=<USER DE PUB DO SEU BROKER>
MQTT_PUB_PASSWORD=<PASSWORD DO PUB DO SEU BROKER>
MQTT_SUB=<USER DE SUB DO SEU BROKER>
MQTT_SUB_PASSWORD=<PASSWORD DO SUB DO SEU BROKER>
```

E depois em terminais diferentes execute o publisher (lembrando que para ambos deve informar o caminho para o arquivo `.env`, e para o publisher é necessário informar o path correto para o csv contendo os dados simulado do sensor) e o subscriber:

- Publisher:

```bash
./run_pub ./.env ./data/dados_sensor_radiacao_solar.csv
```

- Subscriber

```bash
./run_sub ./.env
```

## 3. Estrutura do Projeto

O projeto é composto por dois módulos go, o publisher e o subscriber (ambos podem ser encontrados em suas respectivas pastas).

- `subscriber/main.go`: Arquivo principal que inicia a parte de subscriber do simulador.
- `publisher/main.go`: Arquivo principal que inicia a parte de publisher do simulador.
- `publisher/solar_sensor.go`: Simula a geração de dados para um sensor de radiação solar.
- `data/dados_sensor_radiacao_solar.csv`: Contém dados simulados para o sensor de radiação solar.

## 4. Demonstração do Funcionamento

https://github.com/Lemos1347/inteli-modulo-9-ponderada-3/assets/99190347/d11fb14d-3fdc-4893-ad67-79b48b536722
