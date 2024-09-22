CREATE TABLE "images"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "url"         VARCHAR(255) NOT NULL,
    description   VARCHAR(255),
    "uploaded_at" TIMESTAMP    NOT NULL DEFAULT (now())
);

CREATE TABLE "articles"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "title"      VARCHAR(255) NOT NULL,
    "content"    TEXT         NOT NULL,
    "views"      INT                   DEFAULT 0,
    "min_read"   INT          NOT NULL,
    "created_at" TIMESTAMP    NOT NULL DEFAULT (now()),
    "updated_at" timestamp    NOT NULL DEFAULT (now())
);

CREATE TABLE article_images
(
    article_id BIGINT REFERENCES articles (id) ON DELETE CASCADE,
    image_id   BIGINT REFERENCES images (id) ON DELETE CASCADE,
    PRIMARY KEY (article_id, image_id)
);

CREATE TABLE "article_history"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "article_id"  int,
    "title"       VARCHAR(255),
    "content"     TEXT NOT NULL,
    "modified_at" TIMESTAMP DEFAULT (now())
);

CREATE TABLE "categories"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "name"       VARCHAR(255) UNIQUE NOT NULL,
    "created_at" TIMESTAMP           NOT NULL DEFAULT (now())
);

CREATE TABLE "article_categories"
(
    "article_id"  INT,
    "category_id" INT,
    PRIMARY KEY ("article_id", "category_id")
);


CREATE INDEX ON "articles" ("title");

CREATE INDEX ON "articles" ("views");

CREATE INDEX ON "articles" ("created_at");

CREATE INDEX ON "article_history" ("article_id");

CREATE INDEX ON "categories" ("name");

ALTER TABLE "article_history"
    ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");

ALTER TABLE "article_categories"
    ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");

ALTER TABLE "article_categories"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
