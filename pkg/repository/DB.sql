CREATE TABLE
    News (
        Id bigserial PRIMARY KEY,
        Title text NOT NULL,
        Content text NOT NULL
    );

CREATE TABLE
    NewsCategories (
        NewsId bigint REFERENCES News(Id) ON DELETE CASCADE,
        CategoryId bigint,
        PRIMARY KEY (NewsId, CategoryId)
    );

CREATE TABLE
    Categories(
        Id bigserial PRIMARY KEY,
        Descriprion text NOT NULL
    )