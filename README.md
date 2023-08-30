# Сервис динамического сегментирования пользователей
### Инструкция по запуску:
Файл *compose.yaml* содержит инструкции по контейнерам dynus и postgres. Dockerfile контейнера dynus расположен в корневой папке проекта, а для postgres в *build/postgres*.  
-Для того чтобы начать сборку, необходимо в командной строке, в директории с данным проектом прописать:
*docker compose build --no-cache*
-Для запуска проекта:
*docker compose up -d*
>Возможна ошибка в контейнере dynus при первой инициализации контейнера с БД, повторый запуск dynus решает данную проблему. 

### Требования к входным данным:
-Названия сегментов могут содержать только заглавные и прописные латинские буквы, цифры, а так же нижние подчёркивания;
-Процент пользователей, которые попадут в сегмент при его инициализации, должен быть указан в формате float, причем принадлежать отрезку [0,1];
-Id пользователя - целое положительное;
-Месяц, по которому требуется получить сводку об обновленных записях в формате "YYYY-MM";
-Ttl - строкой в формате временного интервала ("1 year", "1 month" и тд.)
> В противном случае будет получен ответ 400 от сервера с указанием ошибки!

### CURL HTTP запросы к API:
-Создание сегмента:
*curl -XPUT -H"Content-Type: application/json" localhost:8090/slugs \-d '{"name": "<ИМЯ_СЕГМЕНТА>", "chance": "<ПРОЦЕНТ_ПОЛЬЗОВАТЕЛЕЙ_КОТОРЫЕ_БУДУТ_ДОБАВЛЕНЫ_АВТОМАТИЧЕСКИ>"}'*
-Удаление сегмента:
*curl -XDELETE -H"Content-Type: application/json" localhost:8090/slugs \-d '{"name": "<ИМЯ_СЕГМЕНТА>"}'*
-Получение сегментов пользователя:
*curl -XGET -H"Content-Type: application/json" localhost:8090/slugs/<ID_ПОЛЬЗОВАТЕЛЯ>*
-Получение месячной сводки:
*curl -XGET -H"Content-Type: application/json" localhost:8090/slugs/history/<ГОД>-<МЕСЯЦ>*
-Добавление пользователя в сегменты, а так же ttl:
*curl -XPOST -H"Content-Type: application/json" localhost:8090/slugs/<ID_ПОЛЬЗОВАТЕЛЯ> \-d '{"insert_slugs": [<НАЗВАНИЯ_СЕГЕМЕНТОВ_К_ДОБАЛЕНИЮ>, "..."], "delete_slugs": ["<НАЗВАНИЯ_СЕГЕМЕНТОВ_К_УДАЛЕНИЮ>", "..."], "ttl": {"<ИМЕНА_СЕГМЕНТОВ_С_TTL>" : "<TTL>"}}'*
