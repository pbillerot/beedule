-- Adminer 4.8.1 PostgreSQL 14.2 (Debian 14.2-1.pgdg110+1) dump

DROP TABLE IF EXISTS "pluvio";
DROP SEQUENCE IF EXISTS pluvio_id_seq;
CREATE SEQUENCE pluvio_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 START 284 CACHE 1;

CREATE TABLE "public"."pluvio" (
    "id" integer DEFAULT nextval('pluvio_id_seq') NOT NULL,
    "annee" character(4) DEFAULT '2022',
    "mois" integer DEFAULT '3',
    "date" date NOT NULL,
    "mm" integer DEFAULT '0',
    "cumul_mois" integer DEFAULT '0',
    CONSTRAINT "pluvio_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "pluvio" ("id", "annee", "mois", "date", "mm", "cumul_mois") VALUES
(40,	'2021',	5,	'2021-05-15',	12,	112),
(181,	'2020',	5,	'2020-05-01',	6,	69),
(182,	'2020',	5,	'2020-05-02',	9,	69),
(183,	'2020',	5,	'2020-05-03',	1,	69),
(184,	'2020',	5,	'2020-05-04',	3,	69),
(185,	'2020',	5,	'2020-05-10',	34,	69),
(186,	'2020',	5,	'2020-05-08',	4,	69),
(187,	'2020',	5,	'2020-05-11',	5,	69),
(188,	'2020',	5,	'2020-05-12',	2,	69),
(9,	'2021',	1,	'2021-01-24',	3,	156),
(10,	'2021',	1,	'2021-01-27',	11,	156),
(11,	'2021',	1,	'2021-01-28',	30,	156),
(12,	'2021',	1,	'2021-01-29',	12,	156),
(13,	'2021',	1,	'2021-01-30',	13,	156),
(14,	'2021',	1,	'2021-01-31',	19,	156),
(15,	'2021',	2,	'2021-02-01',	26,	114),
(16,	'2021',	2,	'2021-02-02',	24,	114),
(17,	'2021',	2,	'2021-02-05',	8,	114),
(18,	'2021',	2,	'2021-02-06',	11,	114),
(19,	'2021',	2,	'2021-02-11',	10,	114),
(20,	'2021',	2,	'2021-02-09',	17,	114),
(21,	'2021',	2,	'2021-02-12',	8,	114),
(22,	'2021',	2,	'2021-02-18',	6,	114),
(23,	'2021',	2,	'2021-02-22',	2,	114),
(24,	'2021',	3,	'2021-03-10',	6,	41),
(25,	'2021',	3,	'2021-03-11',	12,	41),
(26,	'2021',	3,	'2021-03-12',	3,	41),
(189,	'2020',	5,	'2020-05-23',	5,	69),
(190,	'2020',	6,	'2020-06-04',	3,	105),
(191,	'2020',	6,	'2020-06-03',	3,	105),
(192,	'2020',	6,	'2020-06-05',	4,	105),
(193,	'2020',	6,	'2020-06-09',	2,	105),
(194,	'2020',	6,	'2020-06-10',	3,	105),
(195,	'2020',	6,	'2020-06-11',	13,	105),
(196,	'2020',	6,	'2020-06-12',	4,	105),
(197,	'2020',	6,	'2020-06-13',	11,	105),
(198,	'2020',	6,	'2020-06-14',	11,	105),
(199,	'2020',	6,	'2020-06-15',	4,	105),
(200,	'2020',	6,	'2020-06-16',	23,	105),
(201,	'2020',	6,	'2020-06-17',	5,	105),
(202,	'2020',	6,	'2020-06-18',	5,	105),
(203,	'2020',	6,	'2020-06-21',	4,	105),
(204,	'2020',	6,	'2020-06-27',	2,	105),
(205,	'2020',	6,	'2020-06-28',	8,	105),
(206,	'2020',	7,	'2020-07-25',	4,	13),
(207,	'2020',	7,	'2020-07-26',	6,	13),
(208,	'2020',	7,	'2020-07-05',	3,	13),
(209,	'2020',	8,	'2020-08-12',	1,	41),
(210,	'2020',	8,	'2020-08-10',	3,	41),
(211,	'2020',	8,	'2020-08-13',	19,	41),
(212,	'2020',	8,	'2020-08-15',	7,	41),
(213,	'2020',	8,	'2020-08-20',	1,	41),
(214,	'2020',	8,	'2020-08-21',	1,	41),
(215,	'2020',	8,	'2020-08-22',	4,	41),
(216,	'2020',	8,	'2020-08-27',	3,	41),
(217,	'2020',	8,	'2020-08-30',	2,	41),
(218,	'2020',	9,	'2020-09-18',	5,	74),
(219,	'2020',	9,	'2020-09-19',	4,	74),
(220,	'2020',	9,	'2020-09-21',	19,	74),
(221,	'2020',	9,	'2020-09-23',	12,	74),
(222,	'2020',	9,	'2020-09-24',	16,	74),
(223,	'2020',	9,	'2020-09-25',	2,	74),
(224,	'2020',	9,	'2020-09-26',	1,	74),
(225,	'2020',	9,	'2020-09-30',	15,	74),
(226,	'2020',	10,	'2020-10-01',	25,	173),
(227,	'2020',	10,	'2020-10-02',	8,	173),
(228,	'2020',	10,	'2020-10-03',	36,	173),
(229,	'2020',	10,	'2020-10-04',	7,	173),
(230,	'2020',	10,	'2020-10-05',	11,	173),
(233,	'2020',	10,	'2020-10-09',	3,	173),
(234,	'2020',	10,	'2020-10-11',	2,	173),
(235,	'2020',	10,	'2020-10-12',	2,	173),
(236,	'2020',	10,	'2020-10-13',	2,	173),
(237,	'2020',	10,	'2020-10-20',	3,	173),
(238,	'2020',	10,	'2020-10-23',	3,	173),
(239,	'2020',	10,	'2020-10-24',	8,	173),
(240,	'2020',	10,	'2020-10-25',	21,	173),
(241,	'2020',	10,	'2020-10-26',	8,	173),
(242,	'2020',	10,	'2020-10-27',	17,	173),
(243,	'2020',	10,	'2020-10-28',	8,	173),
(244,	'2020',	10,	'2020-10-31',	4,	173),
(245,	'2020',	11,	'2020-11-01',	8,	41),
(246,	'2020',	11,	'2020-11-02',	4,	41),
(247,	'2020',	11,	'2020-11-07',	5,	41),
(248,	'2020',	11,	'2020-11-11',	3,	41),
(249,	'2020',	11,	'2020-11-15',	15,	41),
(250,	'2020',	11,	'2020-11-03',	1,	41),
(251,	'2020',	11,	'2020-11-18',	5,	41),
(252,	'2020',	12,	'2020-12-03',	29,	202),
(253,	'2020',	12,	'2020-12-05',	8,	202),
(254,	'2020',	12,	'2020-12-07',	9,	202),
(255,	'2020',	12,	'2020-12-10',	18,	202),
(256,	'2020',	12,	'2020-12-09',	9,	202),
(257,	'2020',	12,	'2020-12-11',	8,	202),
(258,	'2020',	12,	'2020-12-14',	13,	202),
(259,	'2020',	12,	'2020-12-15',	4,	202),
(260,	'2020',	12,	'2020-12-16',	3,	202),
(261,	'2020',	12,	'2020-12-17',	2,	202),
(27,	'2021',	3,	'2021-03-13',	10,	41),
(28,	'2021',	3,	'2021-03-18',	2,	41),
(29,	'2021',	3,	'2021-03-25',	1,	41),
(30,	'2021',	3,	'2021-03-26',	2,	41),
(31,	'2021',	4,	'2021-04-10',	6,	15),
(32,	'2021',	4,	'2021-04-11',	9,	15),
(33,	'2021',	5,	'2021-05-04',	14,	112),
(34,	'2021',	5,	'2021-05-06',	16,	112),
(35,	'2021',	5,	'2021-05-09',	16,	112),
(36,	'2021',	5,	'2021-05-10',	5,	112),
(37,	'2021',	5,	'2021-05-11',	2,	112),
(38,	'2021',	5,	'2021-05-12',	5,	112),
(39,	'2021',	5,	'2021-05-13',	3,	112),
(41,	'2021',	5,	'2021-05-16',	7,	112),
(42,	'2021',	5,	'2021-05-17',	3,	112),
(43,	'2021',	5,	'2021-05-18',	2,	112),
(44,	'2021',	5,	'2021-05-19',	1,	112),
(45,	'2021',	5,	'2021-05-22',	6,	112),
(46,	'2021',	5,	'2021-05-23',	12,	112),
(47,	'2021',	5,	'2021-05-24',	8,	112),
(48,	'2021',	6,	'2021-06-01',	2,	119),
(49,	'2021',	6,	'2021-06-02',	3,	119),
(50,	'2021',	6,	'2021-06-03',	16,	119),
(51,	'2021',	6,	'2021-06-04',	3,	119),
(52,	'2021',	6,	'2021-06-16',	8,	119),
(53,	'2021',	6,	'2021-06-17',	14,	119),
(54,	'2021',	6,	'2021-06-18',	11,	119),
(55,	'2021',	6,	'2021-06-20',	7,	119),
(56,	'2021',	6,	'2021-06-21',	10,	119),
(57,	'2021',	6,	'2021-06-22',	1,	119),
(58,	'2021',	6,	'2021-06-26',	3,	119),
(59,	'2021',	6,	'2021-06-27',	10,	119),
(60,	'2021',	6,	'2021-06-28',	13,	119),
(61,	'2021',	6,	'2021-06-29',	18,	119),
(62,	'2021',	7,	'2021-07-03',	3,	49),
(63,	'2021',	7,	'2021-07-04',	8,	49),
(64,	'2021',	7,	'2021-07-05',	2,	49),
(65,	'2021',	7,	'2021-07-06',	2,	49),
(66,	'2021',	7,	'2021-07-12',	20,	49),
(67,	'2021',	7,	'2021-07-13',	5,	49),
(68,	'2021',	7,	'2021-07-24',	9,	49),
(69,	'2021',	8,	'2021-08-02',	6,	42),
(70,	'2021',	8,	'2021-08-03',	6,	42),
(71,	'2021',	8,	'2021-08-04',	4,	42),
(72,	'2021',	8,	'2021-08-05',	7,	42),
(73,	'2021',	8,	'2021-08-06',	1,	42),
(74,	'2021',	8,	'2021-08-07',	15,	42),
(75,	'2021',	8,	'2021-08-08',	2,	42),
(76,	'2021',	8,	'2021-08-10',	1,	42),
(77,	'2021',	9,	'2021-09-09',	9,	45),
(78,	'2021',	9,	'2021-09-10',	4,	45),
(79,	'2021',	9,	'2021-09-14',	1,	45),
(80,	'2021',	9,	'2021-09-17',	3,	45),
(81,	'2021',	9,	'2021-09-25',	7,	45),
(82,	'2021',	9,	'2021-09-26',	8,	45),
(83,	'2021',	9,	'2021-09-28',	7,	45),
(84,	'2021',	10,	'2021-10-02',	27,	54),
(85,	'2021',	10,	'2021-10-03',	5,	54),
(86,	'2021',	10,	'2021-10-04',	7,	54),
(87,	'2021',	10,	'2021-10-20',	4,	54),
(88,	'2021',	10,	'2021-10-21',	2,	54),
(89,	'2021',	10,	'2021-10-25',	2,	54),
(90,	'2021',	10,	'2021-10-29',	7,	54),
(91,	'2021',	11,	'2021-11-01',	5,	36),
(92,	'2021',	11,	'2021-11-02',	9,	36),
(93,	'2021',	11,	'2021-11-03',	13,	36),
(94,	'2021',	11,	'2021-11-07',	2,	36),
(95,	'2021',	11,	'2021-11-13',	2,	36),
(96,	'2021',	11,	'2021-11-25',	2,	36),
(97,	'2021',	11,	'2021-11-28',	3,	36),
(98,	'2021',	12,	'2021-12-01',	10,	177),
(99,	'2021',	12,	'2021-12-03',	29,	177),
(100,	'2021',	12,	'2021-12-04',	5,	177),
(101,	'2021',	12,	'2021-12-05',	2,	177),
(102,	'2021',	12,	'2021-12-06',	4,	177),
(103,	'2021',	12,	'2021-12-07',	24,	177),
(104,	'2021',	12,	'2021-12-10',	13,	177),
(105,	'2021',	12,	'2021-12-12',	2,	177),
(106,	'2021',	12,	'2021-12-23',	2,	177),
(107,	'2021',	12,	'2021-12-25',	7,	177),
(108,	'2021',	12,	'2021-12-26',	16,	177),
(109,	'2021',	12,	'2021-12-27',	33,	177),
(110,	'2021',	12,	'2021-12-28',	16,	177),
(111,	'2021',	12,	'2021-12-29',	14,	177),
(112,	'2021',	9,	'2021-09-03',	6,	45),
(114,	'2022',	1,	'2022-01-04',	4,	59),
(115,	'2022',	1,	'2022-01-06',	3,	59),
(116,	'2022',	1,	'2022-01-08',	18,	59),
(117,	'2022',	1,	'2022-01-09',	13,	59),
(118,	'2022',	1,	'2022-01-16',	3,	59),
(119,	'2022',	1,	'2022-01-19',	2,	59),
(120,	'2022',	1,	'2022-01-30',	2,	59),
(121,	'2022',	2,	'2022-02-04',	5,	59),
(122,	'2022',	2,	'2022-02-06',	4,	59),
(123,	'2022',	2,	'2022-02-13',	5,	59),
(124,	'2020',	1,	'2020-01-01',	11,	151),
(125,	'2020',	1,	'2020-01-03',	5,	151),
(126,	'2020',	1,	'2020-01-07',	5,	151),
(127,	'2020',	1,	'2020-01-09',	15,	151),
(128,	'2020',	1,	'2020-01-12',	7,	151),
(129,	'2020',	1,	'2020-01-13',	2,	151),
(130,	'2020',	1,	'2020-01-14',	4,	151),
(131,	'2020',	1,	'2020-01-16',	9,	151),
(132,	'2020',	1,	'2020-01-17',	8,	151),
(133,	'2020',	1,	'2020-01-18',	2,	151),
(134,	'2020',	1,	'2020-01-27',	25,	151),
(135,	'2020',	1,	'2020-01-26',	7,	151),
(136,	'2020',	1,	'2020-01-30',	26,	151),
(137,	'2020',	1,	'2020-01-28',	7,	151),
(138,	'2020',	1,	'2020-01-29',	8,	151),
(139,	'2020',	1,	'2020-01-31',	10,	151),
(140,	'2020',	2,	'2020-02-01',	24,	109),
(141,	'2020',	2,	'2020-02-02',	10,	109),
(142,	'2020',	2,	'2020-02-03',	3,	109),
(143,	'2020',	2,	'2020-02-08',	2,	109),
(144,	'2020',	2,	'2020-02-09',	4,	109),
(145,	'2020',	2,	'2020-02-11',	2,	109),
(162,	'2020',	3,	'2020-03-09',	5,	130),
(163,	'2020',	3,	'2020-03-10',	8,	130),
(164,	'2020',	3,	'2020-03-11',	4,	130),
(165,	'2020',	3,	'2020-03-15',	10,	130),
(166,	'2020',	3,	'2020-03-16',	5,	130),
(167,	'2021',	3,	'2021-03-18',	5,	41),
(168,	'2020',	3,	'2020-03-04',	30,	130),
(169,	'2020',	3,	'2020-03-05',	14,	130),
(170,	'2020',	3,	'2020-03-06',	9,	130),
(171,	'2020',	4,	'2020-04-06',	6,	60),
(172,	'2020',	4,	'2020-04-17',	4,	60),
(173,	'2020',	4,	'2020-04-18',	1,	60),
(174,	'2020',	4,	'2020-04-19',	10,	60),
(175,	'2020',	4,	'2020-04-20',	1,	60),
(176,	'2020',	4,	'2020-04-25',	13,	60),
(177,	'2020',	4,	'2020-04-27',	2,	60),
(178,	'2020',	4,	'2020-04-28',	2,	60),
(179,	'2020',	4,	'2020-04-29',	8,	60),
(180,	'2020',	4,	'2020-04-30',	13,	60),
(231,	'2020',	10,	'2020-10-06',	4,	173),
(232,	'2020',	10,	'2020-10-07',	1,	173),
(146,	'2020',	2,	'2020-02-10',	1,	109),
(147,	'2020',	2,	'2020-02-13',	5,	109),
(148,	'2020',	2,	'2020-02-16',	15,	109),
(149,	'2020',	2,	'2020-02-17',	2,	109),
(150,	'2020',	2,	'2020-02-18',	5,	109),
(151,	'2020',	2,	'2020-02-23',	2,	109),
(152,	'2020',	2,	'2020-02-24',	7,	109),
(153,	'2020',	2,	'2020-02-26',	9,	109),
(154,	'2020',	2,	'2020-02-27',	9,	109),
(155,	'2021',	2,	'2021-02-28',	2,	114),
(156,	'2020',	2,	'2020-02-29',	7,	109),
(157,	'2020',	2,	'2020-02-28',	2,	109),
(158,	'2020',	3,	'2020-03-01',	21,	130),
(159,	'2020',	3,	'2020-03-02',	5,	130),
(160,	'2020',	3,	'2020-03-03',	14,	130),
(161,	'2020',	3,	'2020-03-08',	5,	130),
(281,	'2022',	3,	'2022-03-02',	1,	3),
(262,	'2020',	12,	'2020-12-19',	13,	202),
(263,	'2020',	12,	'2020-12-20',	4,	202),
(113,	'2022',	1,	'2022-01-03',	14,	59),
(1,	'2021',	1,	'2021-01-12',	5,	156),
(2,	'2021',	1,	'2021-01-13',	4,	156),
(3,	'2021',	1,	'2021-01-14',	5,	156),
(4,	'2021',	1,	'2021-01-15',	10,	156),
(5,	'2021',	1,	'2021-01-20',	8,	156),
(6,	'2021',	1,	'2021-01-21',	18,	156),
(7,	'2021',	1,	'2021-01-22',	15,	156),
(8,	'2021',	1,	'2021-01-23',	3,	156),
(264,	'2020',	12,	'2020-12-21',	23,	202),
(265,	'2020',	12,	'2020-12-22',	2,	202),
(266,	'2020',	12,	'2020-12-23',	8,	202),
(267,	'2020',	12,	'2020-12-26',	3,	202),
(268,	'2020',	12,	'2020-12-27',	35,	202),
(269,	'2020',	12,	'2020-12-28',	5,	202),
(270,	'2020',	12,	'2020-12-29',	6,	202),
(271,	'2022',	2,	'2022-02-14',	2,	59),
(272,	'2022',	2,	'2022-02-15',	19,	59),
(273,	'2022',	2,	'2022-02-16',	8,	59),
(274,	'2022',	2,	'2022-02-17',	4,	59),
(275,	'2022',	2,	'2022-02-18',	2,	59),
(276,	'2022',	2,	'2022-02-19',	2,	59),
(277,	'2022',	2,	'2022-02-20',	5,	59),
(278,	'2022',	2,	'2022-02-21',	1,	59),
(279,	'2022',	2,	'2022-02-22',	1,	59),
(280,	'2022',	2,	'2022-02-22',	1,	59),
(284,	'2022',	3,	'2022-03-04',	2,	3);

-- 2022-03-05 09:04:02.077681+01
