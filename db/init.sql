
create database files;



create table files (
    "name" varchar(255),
    
    "importance" smallint not null, 

    -- 1 high availability (stored in ssd + RAID and cached in memory if less then 20MB) 
    -- 2 high availability (stored in ssd + RAID) protected from data loss
    -- 3 high availability (stored in ssd) no backups
    -- 4 low availability (hard drive with SSD cache)


    "hash" uuid not null,
    "expiration" timestamp not null, -- when initialized the record will be expired in 20 mins
    "size" bigint not null, -- size in bytes
    "encrypted" boolean not null, -- if not the file can be indexed
    "contenttype" varchar(64) -- optional
);