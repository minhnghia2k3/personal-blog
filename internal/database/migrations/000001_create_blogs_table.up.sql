CREATE TABLE "blogs"
(
    "id"         BIGSERIAL PRIMARY KEY,
    "author_id"  INT,
    "title"      VARCHAR(255) NOT NULL,
    "content"    TEXT         NOT NULL,
    "views"      INT                   DEFAULT 0,
    "min_read"   INT          NOT NULL,
    "created_at" TIMESTAMP    NOT NULL DEFAULT (now()),
    "updated_at" timestamp    NOT NULL DEFAULT (now())
);

CREATE TABLE "blog_history"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "blog_id"     int,
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

CREATE TABLE "blog_categories"
(
    "blog_id"     INT,
    "category_id" INT,
    PRIMARY KEY ("blog_id", "category_id")
);


CREATE TABLE "images"
(
    "id"          BIGSERIAL PRIMARY KEY,
    "blog_id"     INT,
    "url"         VARCHAR(255) NOT NULL,
    "uploaded_at" TIMESTAMP    NOT NULL DEFAULT (now())
);


-- INDEXES
CREATE INDEX ON "images" ("blog_id");

CREATE INDEX ON "blogs" ("author_id");

CREATE INDEX ON "blogs" ("title");

CREATE INDEX ON "blogs" ("views");

CREATE INDEX ON "blogs" ("created_at");

CREATE INDEX ON "blog_history" ("blog_id");

CREATE INDEX ON "categories" ("name");

ALTER TABLE "images"
    ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id") ON DELETE CASCADE;

ALTER TABLE "blog_history"
    ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id") ON DELETE CASCADE;

ALTER TABLE "blog_categories"
    ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");

ALTER TABLE "blog_categories"
    ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
