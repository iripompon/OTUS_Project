Проект интеграции системы страхователя с сервисом СЭДО Социального Фонда России (СФР) в части сведений о застрахованном лице.
В системе страхователя - некой организации ведутся сведения о работниках. При приеме сотрудника или любом изменении его перс.данных страхователь обязан передать
сведения в СФР в формате сообщения 86 "Сведения о застрахованном лице". Структура сообщения описана в документе на сайте https://docs-test.fss.ru/sedo.html, документ
"Спецификация сообщений ФК «Реестр получателей услуг»", а также в xsd схеме "Карточка застрахованного лица".
Результат обработки 86 сообщения направляется страхователю в формате 87 сообщения, а сообщением, предназначенным для отправки страхователю информации,
в случае несоответствия сведений о застрахованном с данными в СФР, является 88 сообщение.
Перед отправкой в СФР сообщений они подписываются ЕЦП и шифруются, соответственно отвветные сообщения от СФР расшифровываются. 
Подробно о работе сервиса описано в документе "Спецификация сервис СЭДО v2_0_ЕЦП.docx".
Сервис с шифрованием: https://ecp-test.sfr.gov.ru/sedo/fss/api/soap/SedoGateway?wsdl
Сервис без шифрования: https://ecp-test.sfr.gov.ru/sedo/fss/api/soap/SedoGatewaySimple?wsdl
Необходимо разработать сервис подготовки данных 86 сообщения, подписания ЕЦП и отправки в СФР.
Также необходимо разработать сервис шифрования/расшифрования сообщения в соответствии со спецификацией сервиса СЭДО.
В зависимости от ответа 87 сообщение - успешно принято фондом или 88 сообщение - ошибка в сведениях, полученные от СФР, менять статус и информировать ответственного
за отправку.
На сервере, где будут запускаться сервисы, будет установлено: ПО Крипто Про CSP, закрытые ключи пользователей, выполняющих подписание сообщений, один закрытый ключ служебного сертификата, которым будет
выполняться расшифрование всех входящих сообщений, публичный сертификат СФР, которым будет выполняться шифрование всех исходящих сообщений.
Идентификация закрытого ключа пользователя для подписания сообщения будет выполняться по отпечатку сертификата.
