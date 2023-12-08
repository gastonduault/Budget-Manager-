-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1
-- Généré le : dim. 12 mars 2023 à 11:23
-- Version du serveur : 10.4.27-MariaDB
-- Version de PHP : 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `db_budget_manager`
--

DELIMITER $$
--
-- Procédures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `add_user` (IN `login` VARCHAR(32), IN `pass_word` VARCHAR(32))   BEGIN	
    INSERT INTO user (login, pass_word) VALUES (login, pass_word);
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `change_pass_word` (IN `in_pass_word` VARCHAR(32), IN `in_userID` INT(10))   BEGIN
	UPDATE user
    SET pass_word = in_pass_word
    WHERE userID = in_userID;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `delete_transaction` (IN `in_transactionID` INT(10))   BEGIN
	DECLARE V_accountID INT;
    DECLARE V_amount INT;
    
    SELECT accountID, amount INTO V_accountID, V_amount
    FROM transaction
    WHERE transactionID = in_transactionID;
    
    UPDATE bank_account SET bank_account.balance = bank_account.balance - V_amount WHERE bank_account.accountID = V_accountID;
    DELETE FROM transaction WHERE transactionID = in_transactionID; 
	
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `modify_transaction_amount` (IN `in_transactionID` INT(10), IN `in_amount` INT(9))   BEGIN 
	DECLARE V_accountID INT;
    DECLARE V_amount INT;
    
    SELECT accountID, amount INTO V_accountID, V_amount
    FROM transaction
    WHERE transactionID = in_transactionID;
    
    UPDATE transaction SET amount = in_amount WHERE transactionID = in_transactionID;
    UPDATE bank_account SET bank_account.balance = bank_account.balance - V_amount + in_amount WHERE bank_account.accountID = V_accountID;

END$$

--
-- Fonctions
--
CREATE DEFINER=`root`@`localhost` FUNCTION `connection` (`in_login` VARCHAR(32), `in_pass_word` VARCHAR(32)) RETURNS INT(10)  BEGIN
	DECLARE ret INT;
    SET ret = 0;
    
    SELECT COUNT(*) INTO ret 
    FROM user 
    WHERE login = in_login 
    AND pass_word = in_pass_word;
    
    RETURN ret;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `get_accountID` (`in_login` VARCHAR(32)) RETURNS INT(10)  BEGIN
    DECLARE id INT;
    SET id = -1;
    
    SELECT accountID INTO id
    FROM user INNER JOIN bank_account ON user.userID = bank_account.userID
    WHERE login = in_login;

	RETURN id;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `get_catagory_balance` (`in_categoryID` INT(10)) RETURNS INT(9)  BEGIN 
	DECLARE balance INT;
    
    SELECT SUM(amount) INTO balance
    FROM transaction NATURAL JOIN sub_category
    WHERE sub_category.categoryID = in_categoryID;
    
    RETURN balance;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `get_category_name_from_transaction_ID` (`in_transactionID` INT(10)) RETURNS VARCHAR(32) CHARSET utf8mb4 COLLATE utf8mb4_general_ci  BEGIN 
	DECLARE name VARCHAR(32);
    DECLARE id INT;
    
	SELECT sub_categoryID INTO id
    FROM transaction
    WHERE transactionID = in_transactionID;
    
    SELECT sub_category_name INTO name
    FROM sub_category
    WHERE sub_categoryID = id;
    
    RETURN name;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `get_sub_categoryID` (`in_sub_category_name` VARCHAR(32)) RETURNS INT(10)  BEGIN
	DECLARE id INT;
    SET id = -1;
    
	SELECT sub_categoryID INTO id
    FROM sub_category
    WHERE sub_category_name = in_sub_category_name;
    
    RETURN id;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `get_userID` (`in_login` VARCHAR(32)) RETURNS INT(10)  BEGIN
	DECLARE id INT;
    SET id = -1;
    
    SELECT userID INTO id
    FROM user
    WHERE login = in_login;
    
    RETURN id;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `new_transaction` (`in_login` VARCHAR(32), `in_amount` INT(9), `in_transaction_name` VARCHAR(32), `in_sub_category_name` VARCHAR(32)) RETURNS VARCHAR(64) CHARSET utf8mb4 COLLATE utf8mb4_general_ci  BEGIN
   	DECLARE V_sub_categoryID INT;
    DECLARE V_accountID INT;    

    SET V_sub_categoryID = get_sub_categoryID(in_sub_category_name);

	IF NOT(search_user(in_login)) THEN
        RETURN 'The  user is not in the data base';
    ELSE
    	SET V_accountID = get_accountID(in_login);
    END IF;
    
    IF V_sub_categoryID = -1 THEN
    	RETURN 'The sub category name do not exist';
    END IF;
    
    INSERT INTO transaction (accountID, transaction_name, amount, sub_categoryID) VALUES (V_accountID, in_transaction_name, in_amount, V_sub_categoryID);
    UPDATE bank_account SET bank_account.balance = bank_account.balance + in_amount WHERE bank_account.accountID = V_accountID;

    RETURN 'Transaction made';
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `search_user` (`in_login` VARCHAR(32)) RETURNS TINYINT(1)  BEGIN
    DECLARE find INT;
    SELECT COUNT(*) INTO find
    FROM user
    WHERE login = in_login;

    IF (find > 0) THEN
    	RETURN TRUE;
    ELSE 
    	RETURN FALSE;
    END IF;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Structure de la table `bank_account`
--

CREATE TABLE `bank_account` (
  `accountID` int(11) NOT NULL,
  `userID` int(10) NOT NULL,
  `balance` int(9) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `bank_account`
--

INSERT INTO `bank_account` (`accountID`, `userID`, `balance`) VALUES
(1, 1, 40),
(2, 2, 10000);

-- --------------------------------------------------------

--
-- Structure de la table `category`
--

CREATE TABLE `category` (
  `categoryID` int(10) NOT NULL,
  `category_name` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `category`
--

INSERT INTO `category` (`categoryID`, `category_name`) VALUES
(1, 'Logement');

-- --------------------------------------------------------

--
-- Structure de la table `sub_category`
--

CREATE TABLE `sub_category` (
  `sub_categoryID` int(10) NOT NULL,
  `categoryID` int(10) NOT NULL,
  `sub_category_name` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `sub_category`
--

INSERT INTO `sub_category` (`sub_categoryID`, `categoryID`, `sub_category_name`) VALUES
(1, 1, 'Loyer'),
(4, 1, 'Eau'),
(5, 1, 'Charges');

-- --------------------------------------------------------

--
-- Structure de la table `transaction`
--

CREATE TABLE `transaction` (
  `transactionID` int(10) NOT NULL,
  `accountID` int(10) NOT NULL,
  `transaction_name` varchar(32) NOT NULL,
  `amount` int(9) NOT NULL,
  `sub_categoryID` int(10) NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `transaction`
--

INSERT INTO `transaction` (`transactionID`, `accountID`, `transaction_name`, `amount`, `sub_categoryID`, `date`) VALUES
(1, 1, 'azer', 10, 1, '2023-03-04 15:05:41'),
(4, 1, 'Test', 10, 4, '2023-03-04 15:57:11'),
(5, 1, '', 10, 4, '2023-03-04 16:16:27'),
(6, 1, 'azer', 10, 4, '2023-03-04 16:29:18'),
(7, 1, 'azer', 10, 4, '2023-03-04 16:29:58'),
(8, 1, 'azer', 20, 4, '2023-03-04 16:30:45');

-- --------------------------------------------------------

--
-- Structure de la table `user`
--

CREATE TABLE `user` (
  `userID` int(10) NOT NULL,
  `login` varchar(32) NOT NULL,
  `pass_word` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `user`
--

INSERT INTO `user` (`userID`, `login`, `pass_word`) VALUES
(1, 'julien', 'juju'),
(2, 'Alice', 'Alice'),
(5, 'Corinne', 'Corinne'),
(6, 'Steph', 'Steph'),
(7, 'Steph', 'Steph'),
(8, 'Steph1', 'Steph1'),
(9, 'Julien1', 'Julien1'),
(10, 'Julien2', 'Julien2');

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `bank_account`
--
ALTER TABLE `bank_account`
  ADD PRIMARY KEY (`accountID`),
  ADD KEY `userID` (`userID`);

--
-- Index pour la table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`categoryID`);

--
-- Index pour la table `sub_category`
--
ALTER TABLE `sub_category`
  ADD PRIMARY KEY (`sub_categoryID`),
  ADD KEY `categoryID` (`categoryID`);

--
-- Index pour la table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`transactionID`),
  ADD KEY `fk_accountID` (`accountID`),
  ADD KEY `fk_sub_categoryID` (`sub_categoryID`);

--
-- Index pour la table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`userID`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `bank_account`
--
ALTER TABLE `bank_account`
  MODIFY `accountID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `category`
--
ALTER TABLE `category`
  MODIFY `categoryID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT pour la table `sub_category`
--
ALTER TABLE `sub_category`
  MODIFY `sub_categoryID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT pour la table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `transactionID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT pour la table `user`
--
ALTER TABLE `user`
  MODIFY `userID` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `bank_account`
--
ALTER TABLE `bank_account`
  ADD CONSTRAINT `bank_account_ibfk_1` FOREIGN KEY (`userID`) REFERENCES `user` (`userId`);

--
-- Contraintes pour la table `sub_category`
--
ALTER TABLE `sub_category`
  ADD CONSTRAINT `categoryID` FOREIGN KEY (`categoryID`) REFERENCES `category` (`categoryID`);

--
-- Contraintes pour la table `transaction`
--
ALTER TABLE `transaction`
  ADD CONSTRAINT `fk_accountID` FOREIGN KEY (`accountID`) REFERENCES `bank_account` (`accountID`),
  ADD CONSTRAINT `fk_sub_categoryID` FOREIGN KEY (`sub_categoryID`) REFERENCES `sub_category` (`sub_categoryID`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
