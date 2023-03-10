# whatsapp-api-client-golang

whatsapp-api-client-golang - библиотека на Go, созданная для интеграции с WhatsApp через API
сервиса [BASIS API](https://basis-api.com/). Чтобы начать использовать библиотеку, вам нужно получить ID и token
аккаунта в [личном кабинете](https://ru-cabinet.basis-api.com).

## API

Документация к REST API находится [здесь](https://basis-api.com/docs/api/). Библиотека является оберткой к REST API,
поэтому документация по ссылке выше применима и к самой библиотеке.

## Установка

```shell
go get github.com/Yard-Team/whatsapp-api-client-golang
```

## Импорт

```
import (
	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/api"
)
```

## Авторизация

Чтобы отправить сообщение или выполнить другие методы API, аккаунт WhatsApp в приложении телефона должен быть в
авторизованном состоянии. Для авторизации аккаунта нужно просканировать QR-код
в [личном кабинете](https://ru-cabinet.basis-api.com) с использованием приложения WhatsApp.

## Примеры

### Как инициализировать объект

```
BasisAPI := api.BasisAPI{
    URL: "https://ru-api.basis-api.com",
    IDInstance:       "1234",
    APITokenInstance: "bde035edae3fc00bc116bd112297908d8145e5ba8decc5d884",
}
```

Обратите внимание, что ключи можно получать из переменных среды:

```
IDInstance := os.Getenv("ID_INSTANCE")
APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")
```

### Как создать группу

Ссылка на пример: [createGroup/main.go](examples/createGroup/main.go).

```
response, _ := BasisAPI.Methods().Groups().CreateGroup("groupName", []string{
    "11001234567@c.us",
    "11002345678@c.us",
})
```

### Как отправить вложение

Чтобы отправить вложение, нужно указать первым параметром путь к нужному документу.

Ссылка на пример: [sendFileByUpload/main.go](examples/sendFileByUpload/main.go).

```
response, _ := BasisAPI.Methods().Sending().SendFileByUpload("example.png", map[string]interface{}{
    "chatId": "11001234567@c.us",
})
```

### Как отправить вложение по URI

Ссылка на пример: [sendFileByURL/main.go](examples/sendFileByURL/main.go).

```
response, _ := BasisAPI.Methods().Sending().SendFileByUrl(map[string]interface{}{
    "chatId":   "11001234567@c.us",
    "urlFile":  "https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
    "fileName": "Go-Logo_Blue.svg",
})
```

### Как отправить сообщение

Если у метода API есть необязательные параметры, то в метод библиотеки нужно передавать JSON (`map[string]interface{}`).

Ссылка на пример: [sendMessage/main.go](examples/sendMessage/main.go).

```
response, _ := BasisAPI.Methods().Sending().SendMessage(map[string]interface{}{
    "chatId":  "11001234567@c.us",
    "message": "Any message",
})
```

### Как получать входящие уведомления

Чтобы начать получать уведомления, нужно передать функцию-обработчик в BasisAPIWebhook.Start(). Функция-обработчик
должна содержать 1 параметр (`body map[string]interface{}`). При получении нового уведомления ваша функция-обработчик
будет выполнена. Чтобы перестать получать уведомления, нужно вызвать функцию BasisAPIWebhook.Stop().

Обратите внимание, что нужно импортировать пакет webhook:

```
import (
	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/api"
	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/webhook"
)
```

Ссылка на пример: [webhook/main.go](examples/webhook/main.go).

```
BasisAPIWebhook := webhook.BasisAPIWebhook{
    BasisAPI: BasisAPI,
}

BasisAPIWebhook.Start(func(body map[string]interface{}) {
    typeWebhook := body["typeWebhook"]
    if typeWebhook == "incomingMessageReceived" {
        senderData := body["senderData"]
        chatId := senderData.(map[string]interface{})["chatId"]

        response, _ := BasisAPI.Methods().Sending().SendMessage(map[string]interface{}{
            "chatId":  chatId,
            "message": "Any message",
        })

        BasisAPIWebhook.Stop()
    }
})
```

## Список примеров

| Описание                          | Ссылка на пример                                              |
|-----------------------------------|---------------------------------------------------------------|
| Как создать группу                | [createGroup/main.go](examples/createGroup/main.go)           |
| Как отправить вложение            | [sendFileByUpload/main.go](examples/sendFileByUpload/main.go) |
| Как отправить вложение по URI     | [sendFileByURL/main.go](examples/sendFileByURL/main.go)       |
| Как отправить сообщение           | [sendMessage/main.go](examples/sendMessage/main.go)           |
| Как получать входящие уведомления | [webhook/main.go](examples/webhook/main.go)                   | 

## Список всех методов библиотеки

| Метод API                         | Описание                                                                                                                  | Documentation link                                                                                       |
|-----------------------------------|---------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| `Account().GetSettings`           | Метод предназначен для получения текущих настроек аккаунта                                                                | [GetSettings](https://basis-api.com/docs/api/account/GetSettings/)                                       |
| `Account().SetSettings`           | Метод предназначен для установки настроек аккаунта                                                                        | [SetSettings](https://basis-api.com/docs/api/account/SetSettings/)                                       |
| `Account().GetStateInstance`      | Метод предназначен для получения состояния аккаунта                                                                       | [GetStateInstance](https://basis-api.com/docs/api/account/GetStateInstance/)                             |
| `Account().GetStatusInstance`     | Метод предназначен для получения состояния сокета соединения инстанса аккаунта с WhatsApp                                 | [GetStatusInstance](https://basis-api.com/docs/api/account/GetStatusInstance/)                           |
| `Account().Reboot`                | Метод предназначен для перезапуска аккаунта                                                                               | [Reboot](https://basis-api.com/docs/api/account/Reboot/)                                                 |
| `Account().Logout`                | Метод предназначен для разлогинивания аккаунта                                                                            | [Logout](https://basis-api.com/docs/api/account/Logout/)                                                 |
| `Account().QR`                    | Метод предназначен для получения QR-кода                                                                                  | [QR](https://basis-api.com/docs/api/account/QR/)                                                         |
| `Account().SetProfilePicture`     | Метод предназначен для установки аватара аккаунта                                                                         | [SetProfilePicture](https://basis-api.com/docs/api/account/SetProfilePicture/)                           |
| `Device().GetDeviceInfo`          | Метод предназначен для получения информации об устройстве (телефоне), на котором запущено приложение WhatsApp Business    | [GetDeviceInfo](https://basis-api.com/docs/api/phone/GetDeviceInfo/)                                     |
| `Groups().CreateGroup`            | Метод предназначен для создания группового чата                                                                           | [CreateGroup](https://basis-api.com/docs/api/groups/CreateGroup/)                                        |
| `Groups().UpdateGroupName`        | Метод изменяет наименование группового чата                                                                               | [UpdateGroupName](https://basis-api.com/docs/api/groups/UpdateGroupName/)                                |
| `Groups().GetGroupData`           | Метод получает данные группового чата                                                                                     | [GetGroupData](https://basis-api.com/docs/api/groups/GetGroupData/)                                      |
| `Groups().AddGroupParticipant`    | Метод добавляет участника в групповой чат                                                                                 | [AddGroupParticipant](https://basis-api.com/docs/api/groups/AddGroupParticipant/)                        |
| `Groups().RemoveGroupParticipant` | Метод удаляет участника из группового чата                                                                                | [RemoveGroupParticipant](https://basis-api.com/docs/api/groups/RemoveGroupParticipant/)                  |
| `Groups().SetGroupAdmin`          | Метод назначает участника группового чата администратором                                                                 | [SetGroupAdmin](https://basis-api.com/docs/api/groups/SetGroupAdmin/)                                    |
| `Groups().RemoveAdmin`            | Метод лишает участника прав администрирования группового чата                                                             | [RemoveAdmin](https://basis-api.com/docs/api/groups/RemoveAdmin/)                                        |
| `Groups().SetGroupPicture`        | Метод устанавливает аватар группы                                                                                         | [SetGroupPicture](https://basis-api.com/docs/api/groups/SetGroupPicture/)                                |
| `Groups().LeaveGroup`             | Метод производит выход пользователя текущего аккаунта из группового чата                                                  | [LeaveGroup](https://basis-api.com/docs/api/groups/LeaveGroup/)                                          |
| `Journals().GetChatHistory`       | Метод возвращает историю сообщений чата                                                                                   | [GetChatHistory](https://basis-api.com/docs/api/journals/GetChatHistory/)                                |
| `Journals().GetMessage`           | Метод возвращает сообщение чата                                                                                           | [GetMessage](https://basis-api.com/docs/api/journals/GetMessage/)                                        |
| `Journals().LastIncomingMessages` | Метод возвращает крайние входящие сообщения аккаунта                                                                      | [LastIncomingMessages](https://basis-api.com/docs/api/journals/LastIncomingMessages/)                    |
| `Journals().LastOutgoingMessages` | Метод возвращает крайние отправленные сообщения аккаунта                                                                  | [LastOutgoingMessages](https://basis-api.com/docs/api/journals/LastOutgoingMessages/)                    |
| `Queues().ShowMessagesQueue`      | Метод предназначен для получения списка сообщений, находящихся в очереди на отправку                                      | [ShowMessagesQueue](https://basis-api.com/docs/api/queues/ShowMessagesQueue/)                            |
| `Queues().ClearMessagesQueue`     | Метод предназначен для очистки очереди сообщений на отправку                                                              | [ClearMessagesQueue](https://basis-api.com/docs/api/queues/ClearMessagesQueue/)                          |
| `ReadMark().ReadChat`             | Метод предназначен для отметки сообщений в чате прочитанными                                                              | [ReadChat](https://basis-api.com/docs/api/marks/ReadChat/)                                               |
| `Receiving().ReceiveNotification` | Метод предназначен для получения одного входящего уведомления из очереди уведомлений                                      | [ReceiveNotification](https://basis-api.com/docs/api/receiving/technology-http-api/ReceiveNotification/) |
| `Receiving().DeleteNotification`  | Метод предназначен для удаления входящего уведомления из очереди уведомлений                                              | [DeleteNotification](https://basis-api.com/docs/api/receiving/technology-http-api/DeleteNotification/)   |
| `Receiving().DownloadFile`        | Метод предназначен для скачивания принятых и отправленных файлов                                                          | [DownloadFile](https://basis-api.com/docs/api/receiving/files/DownloadFile/)                             |
| `Sending().SendMessage`           | Метод предназначен для отправки текстового сообщения в личный или групповой чат                                           | [SendMessage](https://basis-api.com/docs/api/sending/SendMessage/)                                       |
| `Sending().SendButtons`           | Метод предназначен для отправки сообщения с кнопками в личный или групповой чат                                           | [SendButtons](https://basis-api.com/docs/api/sending/SendButtons/)                                       |
| `Sending().SendTemplateButtons`   | Метод предназначен для отправки сообщения с интерактивными кнопками из перечня шаблонов в личный или групповой чат        | [SendTemplateButtons](https://basis-api.com/docs/api/sending/SendTemplateButtons/)                       |
| `Sending().SendListMessage`       | Метод предназначен для отправки сообщения с кнопкой выбора из списка значений в личный или групповой чат                  | [SendListMessage](https://basis-api.com/docs/api/sending/SendListMessage/)                               |
| `Sending().SendFileByUpload`      | Метод предназначен для отправки файла, загружаемого через форму (form-data)                                               | [SendFileByUpload](https://basis-api.com/docs/api/sending/SendFileByUpload/)                             |
| `Sending().SendFileByUrl`         | Метод предназначен для отправки файла, загружаемого по ссылке                                                             | [SendFileByUrl](https://basis-api.com/docs/api/sending/SendFileByUrl/)                                   |
| `Sending().SendLocation`          | Метод предназначен для отправки сообщения геолокации                                                                      | [SendLocation](https://basis-api.com/docs/api/sending/SendLocation/)                                     |
| `Sending().SendContact`           | Метод предназначен для отправки сообщения с контактом                                                                     | [SendContact](https://basis-api.com/docs/api/sending/SendContact/)                                       |
| `Sending().SendLink`              | Метод предназначен для отправки сообщения со ссылкой, по которой будут добавлены превью изображения, заголовок и описание | [SendLink](https://basis-api.com/docs/api/sending/SendLink/)                                             |
| `Sending().ForwardMessages`       | Метод предназначен для пересылки сообщений в личный или групповой чат                                                     | [ForwardMessages](https://basis-api.com/docs/api/sending/ForwardMessages/)                               |
| `Service().CheckWhatsapp`         | Метод проверяет наличие аккаунта WhatsApp на номере телефона                                                              | [CheckWhatsapp](https://basis-api.com/docs/api/service/CheckWhatsapp/)                                   |
| `Service().GetAvatar`             | Метод возвращает аватар корреспондента или группового чата                                                                | [GetAvatar](https://basis-api.com/docs/api/service/GetAvatar/)                                           |
| `Service().GetContacts`           | Метод предназначен для получения списка контактов текущего аккаунта                                                       | [GetContacts](https://basis-api.com/docs/api/service/GetContacts/)                                       |
| `Service().GetContactInfo`        | Метод предназначен для получения информации о контакте                                                                    | [GetContactInfo](https://basis-api.com/docs/api/service/GetContactInfo/)                                 |
| `Service().DeleteMessage`         | Метод удаляет сообщение из чата                                                                                           | [DeleteMessage](https://basis-api.com/docs/api/service/deleteMessage/)                                   |
| `Service().ArchiveChat`           | Метод архивирует чат                                                                                                      | [ArchiveChat](https://basis-api.com/docs/api/service/archiveChat/)                                       |
| `Service().UnarchiveChat`         | Метод разархивирует чат                                                                                                   | [UnarchiveChat](https://basis-api.com/docs/api/service/unarchiveChat/)                                   |
| `Service().SetDisappearingChat`   | Метод предназначен для изменения настроек исчезающих сообщений в чатах                                                    | [SetDisappearingChat](https://basis-api.com/docs/api/service/SetDisappearingChat/)                       |
| `BasisAPIWebhook.Start`           | Метод предназначен для старта получения новых уведомлений                                                                 |                                                                                                          |
| `BasisAPIWebhook.Stop`            | Метод предназначен для остановки получения новых уведомлений                                                              |                                                                                                          |

## Лицензия

Лицензия MIT. [LICENSE](LICENSE)
