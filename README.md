# YoFio - Backend Golang - Prueba técnica Carlos André Sánchez M.
### Descripción
En este repositorio se desarrolló una API REST que satisface la solución a la prueba técnica de YoFio.
### Características
* Se desarrolló el endpoint para asignar créditos
* Se desarrolló el endpoint para obtener estadísticas de los créditos
* Se creó una BD para el guardado de asignaciones de crédito
### API Endpoints

### URL de la API para consumirla: [3.133.119.213:8000](3.133.119.213:8000)

### POST - /credit-assignment
Ruta para calcular la asignación de créditos

`3.133.119.213:8000/credit-assignment`
#### Request
```json
  {
      "investment":3000
  }
```
|Parametro|Tipo|Descripción|
|---|---|---|
|investment|Entero|El monto de inversión deseado

### Response

### Success Response
```json 
 {
    "credit_type_300": 3,
    "credit_type_500": 0,
    "credit_type_700": 3
 }
```
|Parametro|Tipo|Descripción|
|---|---|---|
|credit_type_300|Entero|Cantidad créditos por montos de 300
|credit_type_500|Entero|Cantidad créditos por montos de 500
|credit_type_700|Entero|Cantidad créditos por montos de 700

### Error Response
```json 
 {
    "error": "no se pudo realizar la inversion"
 }
```

---

### POST - /statistics
Ruta para obtener las estadísticas de asignación de créditos

`3.133.119.213:8000/statistics`

### Response

### Success Response
```json 
  {
    "total_assignations": 7,
    "total_success_assignation": 4,
    "total_failed_assignation": 3,
    "average_success_assignation": 3200,
    "average_failed_assignation": 167.3333
  }
```
|Parametro|Tipo|Descripción|
|---|---|---|
|total_assignations|Entero|Total de asignaciones realizadas
|total_success_assignation|Entero|Total de asignaciones exitosas
|total_failed_assignation|Entero|Total de asignaciones no exitosas
|average_success_assignation|Flotante|Promedio de inversión exitosa
|average_failed_assignation|Flotante|Promedio de inversión no exitosa

### Error Response
```json 
 {
    "error": "no se pudieron obtener los datos"
 }
```




