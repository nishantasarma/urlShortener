cat >> sql.sh << EOF

create database urlmap;
use urlmap;
create table urls(

    shortUrl varchar(4) NOT NULL,
    longUrl varchar(30) NOT NULL
);

insert into urls values("abcd", "youtube.com");

EOF