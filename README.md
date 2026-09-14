Project_template

# Задание 1. Анализ и планирование

## 1. Описание функциональности монолитного приложения
No real evidence on functionality for Heat system

possibly, system manages both Temperature monitoring/Heat system using same entity (possible heat management candidate - PATCH api/v1/sensors/{id}/value)

Heat system control:
Users can:
- Control termal settings (no code)

System supports:
- Set desired temperature over api/v1/sensors/{id}/value?

Temperature monitoring:
Users can:
- check thermal sensor readings
System supports:
- Call Health check
- Get all sensors 
- Get a specific sensor
- Create a new sensor
- Update a sensor
- Delete a sensor
- Update a sensor's value and status

…
## 2. Анализ архитектуры монолитного приложения

Tecnical stack is legit and consists of:
- Database: Postgre SQL (v.16) 
- Backend: one POC application written on Go 1.22 (can be updated to 1.27+).
- Both services are designed to run in containers, 

Implementation described as monolityc, call implementation used syncronus restful API

All calls initiated from server to sensors

System doccumentation claims control of one house sensors by installed instance.
From the other side, System operation schema is contraversal and claims different: system is designed utilizing single instance of service/database to cover all the customers (inside of one connected area).
No any distinct smart house level separation found in db/code of current state

No auntentifiaction or authorization implemented

User and administrator enproints are part of one root level API

Based on current servicing scenario sensors installation and configuration is done inplace, by authorized person.

No UI (web/mobile-app) referenced at all. Not clear, how the user operates the system

Current functionality is minimal, system is in the middle of development, supplied code has no thermal sensor reads functioanlity at all (documentation claims direct per sensor calls, no such BL or database configuration schema sexist).

Heat control is implemented with one single endpoint: 
PATCH api/v1/sensors/1/value

No separation between termal sensor and heat system controls exist

## 3. Определение доменов и границы контекстов
In current state there are two domains in the system I could identify:
- WarmHouse Management
- Sensors usage

## 4. Проблемы монолитного решения
- direct server-to-sensor pull model
- in-place deployment (whole system installed per each house)
- remote service and customer support not available
- one payment installation fee usually high because of hardware costs
- no flexibility on the level of extension / no HA and disaster recovery
- all errored calls are missed (no async/QOS)

## 5. Визуализация контекста системы — диаграмма С4
[Curent state C4 Context](docs/architecture/curent_state/c4_context.puml)

# Задание 2. Проектирование микросервисной архитектуры

## Диаграмма контейнеров (Containers)
[Microservices C4 Containers](docs/architecture/microservices/c4_containers.puml)

## Диаграмма компонентов (Components)
[C4 Smart-house Components Diagram](docs/architecture/microservices/c4_components_smartHouse.puml)

## Диаграмма кода (Code)
[C4 Code - Sensors Management](docs/architecture/microservices/c4_code_sensorManagement.puml)
[C4 Code - Heating](docs/architecture/microservices/c4_code_Heating.puml)

# Задание 3. Разработка ER-диаграммы
[ER - main flows](docs/architecture/microservices/er.puml)

# Задание 4. Создание и документирование API
## 1. Тип API
Two types of API are going to be used:
- RESTful API UI <> API gateway
- async Microservices(through Plugin) <> MQTT router

## 2. Документация API
 ### API (REST), (AI-assisted)
 **Sensors API (REST)**: [docs/swagger/sensors.yaml](docs/API/swagger/sensors.yaml) 
 **Heating API (REST)**: [docs/swagger/heating.yaml](docs/API/swagger/heating.yaml) 

 ### AsyncAPI (MQTT), (AI-assisted)
*Asynchronous communication between the Plugin layer, MQTT Router, and smart devices.*
* **Telemetry Streams**: `telemetry/sensors/{sensor_id}`
  * **Description**: Real-time updates from sensors (temperature, humidity, motion).
  * **Payload**: `{"value": 22.5, "unit": "C", "timestamp": "..."}`
* **Command Dispatch**: `commands/heating/{sensor_id}`
  * **Description**: Instructions sent from the API Gateway to heating actuators.
  * **Payload**: `{"action": "SET_TEMP", "target": 24.0}`
* **System Alerts**: `alerts/security`
  * **Description**: High-priority notifications (e.g., unauthorized entry, sensor failure).
  * **Payload**: `{"alert_type": "MOTION_DETECTED", "location": "Kitchen"}`

# Задание 5. Работа с docker и docker-compose
Minimal implementation of temperature_api has been creted.

All logic mentioned is implemented as part of the new service

Docker-compose updated initially to host the existing service with db initi on db conainer creation (no db update style for now)

Added temperature api_docker config + file that connects to the shared network 

New service runs on port 8081 as expected

Postman checks produced for:
Create Sensor
Get All Sensors

Each call brings random temperature values