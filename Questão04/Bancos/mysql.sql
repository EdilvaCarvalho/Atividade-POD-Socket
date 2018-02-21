CREATE TABLE `pessoa` (
        `uid` INT(10) NOT NULL AUTO_INCREMENT,
        `nome` VARCHAR(64) NULL DEFAULT NULL,
        `curso` VARCHAR(64) NULL DEFAULT NULL,
        `cidade` VARCHAR(64) NULL DEFAULT NULL,
        PRIMARY KEY (`uid`)
    );
    
    
select * from pessoa