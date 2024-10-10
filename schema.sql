
CREATE TABLE "folder" (
    "id"  SERIAL  NOT NULL,
    "chat_id" begserial   NOT NULL,
    "name" text   NOT NULL,
    CONSTRAINT "pk_folder" PRIMARY KEY (
        "id"
     )
);

CREATE TABLE "files" (
    "id"  SERIAL  NOT NULL,
    "folder_id" int   NOT NULL,
    "location" text   NOT NULL,
    CONSTRAINT "pk_files" PRIMARY KEY (
        "id"
     )
);

ALTER TABLE "files" ADD CONSTRAINT "fk_files_folder_id" FOREIGN KEY("folder_id")
REFERENCES "folder" ("id");