# lABoratory
Platform for creating and using A/B testing experiments.

## Description of the project and objective to be achieved
Las pruebas A/B son un tipo de experimentos que se utilizan en el contexto del marketing digital y la analítica web para identificar cambios que maximizan un resultado determinado. Por ejemplo, clientes diferentes de una página de compras online pueden ver productos diferentes, anuncios en diferentes lugares, o precios diferentes para el mismo producto, basados en dicho experimento. Mediante la obtención de métricas vinculadas al experimento se puede determinar la optimalidad de cada opción y finalmente tomar una decisión en base a los resultados.

Este Trabajo Fin de Grado consiste en el desarrollo de una plataforma para la creación de experimentos A/B. El usuario podrá configurar mediante un interfaz web un experimento, con un nombre determinado (clave única), así como un conjunto de activaciones, que serán normalmente el tratamiento C (grupo de control) y la activación A1 (experimento activo), pero que pueden ser más (A2, A3, etc.). Asimismo, podrá asignar determinados porcentajes de asignación a las activaciones, totalizando el 100% (por ejemplo 50% al grupo de control, 25% a la activación A1 y 25% a la activación A2). Opcionalmente se proporcionará la funcionalidad para deshabilitar el experimento totalmente (100% de asignación a C) o para lanzar completamente alguna de las activaciones (100% a A1, A2 o la activación deseada).

Por otro lado, un usuario de una web que a su vez utilice este servicio de experimentación siempre debería obtener el mismo tratamiento. Por lo tanto, el TFG incluirá un servicio web que expondrá una API que permitirá obtener el valor (C, A1, A2, etc) del experimento para una determinada clave única: user ID, customer ID (la semántica debe de ser abstracta). Cuando se consulte el tratamiento para un experimento por primera vez para una clave, se asignará un valor aleatorio en base a los porcentajes fijados para él. Sin embargo, posteriores consultas para la misma clave deben devolver el mismo valor para garantizar coherencia del experimento de cara al usuario final. Cuando se cambie la asignación de las activaciones se intentará mantener la asignación en la medida de lo posible para las parejas clave/tratamiento ya generadas, respetando no obstante los nuevos porcentajes. Finalmente, debe ser posible crear una excepción (override) para una determinada clave, asignando un tratamiento específico.

Se valorará la utilización de tecnologías Cloud para la creación del servicio web en aras de garantizar la escalabilidad y redundancia de éste.


## Task to do
 * Investigar las tecnologías que se usarán.
 * Desarrollar el análisis, diseño, implementación, pruebas y puesta en funcionamiento de la aplicación.

## Techonologies to use
 - [ ] .Net
 - [x] Go
 - [ ] Java
 - [x] Angular
 - [ ] React

## Key words
`Web services, Web development, Cloud, A/B testing`


