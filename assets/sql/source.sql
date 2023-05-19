create database wkg;

# drop database wkg;

use wkg;

# INSERT INTO `websites` (`cid`,`favicon`,`website`,`title`,`headers`,`finger`,`screenshot`,`updateTime`) VALUES (5,'1558343002','http://mock.coding.bjf.yun.unionpay.com:80','CODING - 一站式�','','','','2021-11-18 10:50:21');

# ALTER DATABASE wkg CHARACTER SET gbk COLLATE gbk;

# drop table users;

create table
    users (
        id int primary key not null auto_increment,
        username varchar(20),
        password varchar(32),
        createTime text
    ) default charset = gbk,
    engine = innodb;

insert into users values (1, 'gelen', 'gelen');

ALTER TABLE users ADD COLUMN createTime text NOT NULL;

# drop table company;

create table
    company (
        id int primary key not null auto_increment,
        projectType varchar(10) not null default '-',
        #项目类型，第三方，公益，CNVD，SRC
        companyName text,
        domain text,
        #待搜集的域名信息
        keyWord varchar(500) not null default '-',
        srcUrl varchar(50) not null default '-',
        #SRC网址
        monitorStatus bool not null default true,
        monitorRate int not null default 24,
        remarks text,
        vulnScanStatus bool not null default false,
        vulnScanRate int not null default 24,
        #以小时为单位
        lastUpdateTime varchar(20) not null default '-'
    ) default charset = gbk;

# drop table domains;

create table
    domains (
        id int primary key auto_increment,
        cid int not null,
        domain text,
        type varchar(10) not null default '-',
        source text,
        remarks text,
        updateTime varchar(20) not null default '-',
        isNew bool default true
    ) default charset = gbk;

# drop table websites;

create table
    websites (
        id int primary key auto_increment,
        cid int not null,
        domain varchar(150),
        ips varchar(300),
        website varchar(150) not null default '-',
        favicon varchar(32) default '-',
        #favicon.
        faviconUrl text,
        #favicon
        title text,
        CDN bool,
        headers text,
        cert text,
        finger varchar(500) not null default '-',
        screenshot longblob,
        updateTime varchar(20) not null default '-',
        remarks text,
        isNew bool default true
    ) default charset = gbk;

use wkg;

create table
    notice (
        id int primary key auto_increment,
        cid int not null,
        title varchar(256) not null,
        #ipv6最长46位
        content text not null,
        updateTime varchar(20) not null
    ) default charset = gbk;

drop table ips;

create table
    ips (
        id int primary key auto_increment,
        cid int not null,
        ip varchar(46) not null default '-',
        #ipv6最长46位
        os varchar(50) not null default '-',
        indomains varchar(3000) not null default '-',
        geo varchar(30) not null default '-',
        remarks text,
        updateTime varchar(20) not null default '-',
        isNew bool default true
    ) default charset = gbk;

# drop table services;

create table
    services (
        id int primary key auto_increment,
        cid int not null,
        service varchar(100) not null default '-',
        host text not null,
        port varchar(500) not null default '-',
        product varchar(500) not null default '-',
        updateTime varchar(20) not null default '-',
        remarks text,
        isNew bool default true
    ) default charset = gbk;

# drop table apps;

create table
    apps (
        id int primary key auto_increment,
        cid int not null,
        appname text not null,
        remarks text,
        updateTime varchar(20) not null default '-',
        isNew bool default true
    ) default charset = gbk;

# drop table webchatOfficeAccount;

create table
    webchatofficeaccount (
        id int primary key auto_increment,
        cid int not null,
        name text not null,
        remarks text not null,
        updateTime varchar(20) not null default '-',
        isNew bool default true
    ) default charset = gbk;

# drop table miniProgram;

create table
    miniprogram (
        id int primary key auto_increment,
        cid int not null,
        name text,
        remarks text,
        updateTime varchar(20) not null default '-',
        isNew bool default true
    ) default charset = gbk;

# drop table systemConfig;

create table
    systemConfig (
        id int primary key,
        emailNotifyEnable bool default false,
        emailServerAddr varchar(30),
        emailUserName varchar(30),
        emailPassword varchar(30),
        weChatNotify_Enable bool default false,
        weChatKey varchar(45),
        dingtalkNotfyEnable bool default false,
        updateTime varchar(20) not null default '-',
        dingtalkAccessToken varchar(80)
    ) default charset = gbk;

# drop table knowledge;

create table
    knowledge (
        id int primary key auto_increment,
        title text,
        content text,
        ckey varchar(33),
        shared bool,
        hit int,
        author varchar(20),
        tags varchar(30) not null;

createTime text,
updateTime text,
parentId int
) default charset = gbk;

ALTER TABLE knowledge ADD COLUMN author varchar(20) NOT NULL;

ALTER TABLE knowledge ADD COLUMN createTime text NOT NULL;

# drop table category;

create table
    category (
        id int primary key auto_increment,
        parentId int,
        title text,
        isLeaf bool,
        ckey varchar(33)
    ) default charset = gbk;

# drop table blackiplist;

create table
    blackiplist (
        id int primary key auto_increment,
        ip tinytext,
        count int
    ) default charset = gbk;

# drop table bugReport;

create table
    bugReport (
        id int primary key auto_increment,
        detail text,
        updateTime tinytext,
        isNew bool
    ) default charset = gbk;

# drop table bugCollect;

create table
    bugCollect (
        id int primary key auto_increment,
        detail text,
        updateTime tinytext
    ) default charset = gbk;

drop table agent;

create table
    agent (
        agent_id VARCHAR(50) primary key,
        IntranetIpv4 TEXT,
        IntranetIpv6 TEXT,
        ExtranetIpv4 TEXT,
        ExtranetIpv6 TEXT,
        host_name TEXT,
        version TEXT,
        product TEXT,
        kernel_version TEXT,
        arch TEXT,
        platform TEXT,
        platform_family TEXT,
        platform_version TEXT,
        cpu TEXT,
        rss TEXT,
        read_speed TEXT,
        write_speed TEXT,
        pid TEXT,
        nfd TEXT,
        start_time TEXT,
        total_mem TEXT,
        host_serial TEXT,
        host_id TEXT,
        host_model TEXT,
        host_vendor TEXT,
        cpu_name TEXT,
        boot_time TEXT,
        cpu_usage TEXT,
        mem_usage TEXT,
        update_time TEXT
    ) default charset = gbk;

drop table task;

create table
    task (
        taskId varchar(256) primary key,
        taskStatus bool,
        taskName varchar(30),
        taskType varchar(5),
        rate int,
        companyRange int,
        companyId text,
        comanyName text,
        pocId int,
        pocName text,
        assetRange int,
        DomainCollectionType int,
        IPCollectionType int,
        portRange int,
        updateTime text
    ) default charset = gbk;

drop table poc;

create table
    poc (
        id int primary key auto_increment,
        pocName varchar(100),
        pocContent text,
        updateTime tinytext
    ) default charset = gbk;