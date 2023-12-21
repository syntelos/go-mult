/*
 * MULTI/FORMAT/CODEC I/O
 * Copyright 2023 John Douglas Pritchard, Syntelos
 *
 *
 * References
 * 
 * https://github.com/multiformats/multicodec/blob/master/table.csv
 * https://multiformats.io/
 */
package mult
/*
 * Multicodec header as VARINT/ULEB64.
 */
type Header uint64
/*
 * Multicodec identifier as GOPL ID.
 */
type Identifier string
/*
 * Multicodec constants include GOPL domain UNKNOWNs.
 */
const (
	HeaderUNKNOWN		Header = 0
	HeaderMultihashIDENTITY	Header	=	0x00
	HeaderCidCIDV1	Header	=	0x01
	HeaderCidCIDV2	Header	=	0x02
	HeaderCidCIDV3	Header	=	0x03
	HeaderMultiaddrIP4	Header	=	0x04
	HeaderMultiaddrTCP	Header	=	0x06
	HeaderMultihashSHA1	Header	=	0x11
	HeaderMultihashSHA2256	Header	=	0x12
	HeaderMultihashSHA2512	Header	=	0x13
	HeaderMultihashSHA3512	Header	=	0x14
	HeaderMultihashSHA3384	Header	=	0x15
	HeaderMultihashSHA3256	Header	=	0x16
	HeaderMultihashSHA3224	Header	=	0x17
	HeaderMultihashSHAKE128	Header	=	0x18
	HeaderMultihashSHAKE256	Header	=	0x19
	HeaderMultihashKECCAK224	Header	=	0x1a
	HeaderMultihashKECCAK256	Header	=	0x1b
	HeaderMultihashKECCAK384	Header	=	0x1c
	HeaderMultihashKECCAK512	Header	=	0x1d
	HeaderMultihashBLAKE3	Header	=	0x1e
	HeaderMultihashSHA2384	Header	=	0x20
	HeaderMultiaddrDCCP	Header	=	0x21
	HeaderHashMURMUR3X6464	Header	=	0x22
	HeaderHashMURMUR332	Header	=	0x23
	HeaderMultiaddrIP6	Header	=	0x29
	HeaderMultiaddrIP6ZONE	Header	=	0x2a
	HeaderMultiaddrIPCIDR	Header	=	0x2b
	HeaderNamespacePATH	Header	=	0x2f
	HeaderMultiformatMULTICODEC	Header	=	0x30
	HeaderMultiformatMULTIHASH	Header	=	0x31
	HeaderMultiformatMULTIADDR	Header	=	0x32
	HeaderMultiformatMULTIBASE	Header	=	0x33
	HeaderMultiformatVARSIG	Header	=	0x34
	HeaderMultiaddrDNS	Header	=	0x35
	HeaderMultiaddrDNS4	Header	=	0x36
	HeaderMultiaddrDNS6	Header	=	0x37
	HeaderMultiaddrDNSADDR	Header	=	0x38
	HeaderSerializationPROTOBUF	Header	=	0x50
	HeaderIpldCBOR	Header	=	0x51
	HeaderIpldRAW	Header	=	0x55
	HeaderMultihashDBLSHA2256	Header	=	0x56
	HeaderSerializationRLP	Header	=	0x60
	HeaderSerializationBENCODE	Header	=	0x63
	HeaderIpldDAGPB	Header	=	0x70
	HeaderIpldDAGCBOR	Header	=	0x71
	HeaderIpldLIBP2PKEY	Header	=	0x72
	HeaderIpldGITRAW	Header	=	0x78
	HeaderIpldTORRENTINFO	Header	=	0x7b
	HeaderIpldTORRENTFILE	Header	=	0x7c
	HeaderIpldLEOFCOINBLOCK	Header	=	0x81
	HeaderIpldLEOFCOINTX	Header	=	0x82
	HeaderIpldLEOFCOINPR	Header	=	0x83
	HeaderMultiaddrSCTP	Header	=	0x84
	HeaderIpldDAGJOSE	Header	=	0x85
	HeaderIpldDAGCOSE	Header	=	0x86
	HeaderNamespaceLBRY	Header	=	0x8c
	HeaderIpldETHBLOCK	Header	=	0x90
	HeaderIpldETHBLOCKLIST	Header	=	0x91
	HeaderIpldETHTXTRIE	Header	=	0x92
	HeaderIpldETHTX	Header	=	0x93
	HeaderIpldETHTXRECEIPTTRIE	Header	=	0x94
	HeaderIpldETHTXRECEIPT	Header	=	0x95
	HeaderIpldETHSTATETRIE	Header	=	0x96
	HeaderIpldETHACCOUNTSNAPSHOT	Header	=	0x97
	HeaderIpldETHSTORAGETRIE	Header	=	0x98
	HeaderIpldETHRECEIPTLOGTRIE	Header	=	0x99
	HeaderIpldETHRECEIPTLOG	Header	=	0x9a
	HeaderKeyAES128	Header	=	0xa0
	HeaderKeyAES192	Header	=	0xa1
	HeaderKeyAES256	Header	=	0xa2
	HeaderKeyCHACHA128	Header	=	0xa3
	HeaderKeyCHACHA256	Header	=	0xa4
	HeaderIpldBITCOINBLOCK	Header	=	0xb0
	HeaderIpldBITCOINTX	Header	=	0xb1
	HeaderIpldBITCOINWITNESSCOMMITMENT	Header	=	0xb2
	HeaderIpldZCASHBLOCK	Header	=	0xc0
	HeaderIpldZCASHTX	Header	=	0xc1
	HeaderMultiformatCAIP50	Header	=	0xca
	HeaderNamespaceSTREAMID	Header	=	0xce
	HeaderIpldSTELLARBLOCK	Header	=	0xd0
	HeaderIpldSTELLARTX	Header	=	0xd1
	HeaderMultihashMD4	Header	=	0xd4
	HeaderMultihashMD5	Header	=	0xd5
	HeaderIpldDECREDBLOCK	Header	=	0xe0
	HeaderIpldDECREDTX	Header	=	0xe1
	HeaderNamespaceIPLD	Header	=	0xe2
	HeaderNamespaceIPFS	Header	=	0xe3
	HeaderNamespaceSWARM	Header	=	0xe4
	HeaderNamespaceIPNS	Header	=	0xe5
	HeaderNamespaceZERONET	Header	=	0xe6
	HeaderKeySECP256K1PUB	Header	=	0xe7
	HeaderNamespaceDNSLINK	Header	=	0xe8
	HeaderKeyBLS12381G1PUB	Header	=	0xea
	HeaderKeyBLS12381G2PUB	Header	=	0xeb
	HeaderKeyX25519PUB	Header	=	0xec
	HeaderKeyED25519PUB	Header	=	0xed
	HeaderKeyBLS12381G1G2PUB	Header	=	0xee
	HeaderKeySR25519PUB	Header	=	0xef
	HeaderIpldDASHBLOCK	Header	=	0xf0
	HeaderIpldDASHTX	Header	=	0xf1
	HeaderIpldSWARMMANIFEST	Header	=	0xfa
	HeaderIpldSWARMFEED	Header	=	0xfb
	HeaderIpldBEESON	Header	=	0xfc
	HeaderMultiaddrUDP	Header	=	0x0111
	HeaderMultiaddrP2PWEBRTCSTAR	Header	=	0x0113
	HeaderMultiaddrP2PWEBRTCDIRECT	Header	=	0x0114
	HeaderMultiaddrP2PSTARDUST	Header	=	0x0115
	HeaderMultiaddrWEBRTCDIRECT	Header	=	0x0118
	HeaderMultiaddrWEBRTC	Header	=	0x0119
	HeaderMultiaddrP2PCIRCUIT	Header	=	0x0122
	HeaderIpldDAGJSON	Header	=	0x0129
	HeaderMultiaddrUDT	Header	=	0x012d
	HeaderMultiaddrUTP	Header	=	0x012e
	HeaderHashCRC32	Header	=	0x0132
	HeaderHashCRC64ECMA	Header	=	0x0164
	HeaderMultiaddrUNIX	Header	=	0x0190
	HeaderMultiaddrTHREAD	Header	=	0x0196
	HeaderMultiaddrP2P	Header	=	0x01a5
	HeaderMultiaddrHTTPS	Header	=	0x01bb
	HeaderMultiaddrONION	Header	=	0x01bc
	HeaderMultiaddrONION3	Header	=	0x01bd
	HeaderMultiaddrGARLIC64	Header	=	0x01be
	HeaderMultiaddrGARLIC32	Header	=	0x01bf
	HeaderMultiaddrTLS	Header	=	0x01c0
	HeaderMultiaddrSNI	Header	=	0x01c1
	HeaderMultiaddrNOISE	Header	=	0x01c6
	HeaderMultiaddrSHS	Header	=	0x01c8
	HeaderMultiaddrQUIC	Header	=	0x01cc
	HeaderMultiaddrQUICV1	Header	=	0x01cd
	HeaderMultiaddrWEBTRANSPORT	Header	=	0x01d1
	HeaderMultiaddrCERTHASH	Header	=	0x01d2
	HeaderMultiaddrWS	Header	=	0x01dd
	HeaderMultiaddrWSS	Header	=	0x01de
	HeaderMultiaddrP2PWEBSOCKETSTAR	Header	=	0x01df
	HeaderMultiaddrHTTP	Header	=	0x01e0
	HeaderIpldSWHID1SNP	Header	=	0x01f0
	HeaderIpldJSON	Header	=	0x0200
	HeaderSerializationMESSAGEPACK	Header	=	0x0201
	HeaderSerializationCAR	Header	=	0x0202
	HeaderSerializationIPNSRECORD	Header	=	0x0300
	HeaderLibp2pLIBP2PPEERRECORD	Header	=	0x0301
	HeaderLibp2pLIBP2PRELAYRSVP	Header	=	0x0302
	HeaderLibp2pMEMORYTRANSPORT	Header	=	0x0309
	HeaderSerializationCARINDEXSORTED	Header	=	0x0400
	HeaderSerializationCARMULTIHASHINDEXSORTED	Header	=	0x0401
	HeaderTransportTRANSPORTBITSWAP	Header	=	0x0900
	HeaderTransportTRANSPORTGRAPHSYNCFILECOINV1	Header	=	0x0910
	HeaderTransportTRANSPORTIPFSGATEWAYHTTP	Header	=	0x0920
	HeaderMultiformatMULTIDID	Header	=	0x0d1d
	HeaderMultihashSHA2256TRUNC254PADDED	Header	=	0x1012
	HeaderMultihashSHA2224	Header	=	0x1013
	HeaderMultihashSHA2512224	Header	=	0x1014
	HeaderMultihashSHA2512256	Header	=	0x1015
	HeaderHashMURMUR3X64128	Header	=	0x1022
	HeaderMultihashRIPEMD128	Header	=	0x1052
	HeaderMultihashRIPEMD160	Header	=	0x1053
	HeaderMultihashRIPEMD256	Header	=	0x1054
	HeaderMultihashRIPEMD320	Header	=	0x1055
	HeaderMultihashX11	Header	=	0x1100
	HeaderKeyP256PUB	Header	=	0x1200
	HeaderKeyP384PUB	Header	=	0x1201
	HeaderKeyP521PUB	Header	=	0x1202
	HeaderKeyED448PUB	Header	=	0x1203
	HeaderKeyX448PUB	Header	=	0x1204
	HeaderKeyRSAPUB	Header	=	0x1205
	HeaderKeySM2PUB	Header	=	0x1206
	HeaderKeyED25519PRIV	Header	=	0x1300
	HeaderKeySECP256K1PRIV	Header	=	0x1301
	HeaderKeyX25519PRIV	Header	=	0x1302
	HeaderKeySR25519PRIV	Header	=	0x1303
	HeaderKeyRSAPRIV	Header	=	0x1305
	HeaderKeyP256PRIV	Header	=	0x1306
	HeaderKeyP384PRIV	Header	=	0x1307
	HeaderKeyP521PRIV	Header	=	0x1308
	HeaderMultihashKANGAROOTWELVE	Header	=	0x1d01
	HeaderEncryptionAESGCM256	Header	=	0x2000
	HeaderMultiaddrSILVERPINE	Header	=	0x3f42
	HeaderMultihashSM3256	Header	=	0x534d
	HeaderHashSHA256A	Header	=	0x7012
	HeaderMultihashBLAKE2B8	Header	=	0xb201
	HeaderMultihashBLAKE2B16	Header	=	0xb202
	HeaderMultihashBLAKE2B24	Header	=	0xb203
	HeaderMultihashBLAKE2B32	Header	=	0xb204
	HeaderMultihashBLAKE2B40	Header	=	0xb205
	HeaderMultihashBLAKE2B48	Header	=	0xb206
	HeaderMultihashBLAKE2B56	Header	=	0xb207
	HeaderMultihashBLAKE2B64	Header	=	0xb208
	HeaderMultihashBLAKE2B72	Header	=	0xb209
	HeaderMultihashBLAKE2B80	Header	=	0xb20a
	HeaderMultihashBLAKE2B88	Header	=	0xb20b
	HeaderMultihashBLAKE2B96	Header	=	0xb20c
	HeaderMultihashBLAKE2B104	Header	=	0xb20d
	HeaderMultihashBLAKE2B112	Header	=	0xb20e
	HeaderMultihashBLAKE2B120	Header	=	0xb20f
	HeaderMultihashBLAKE2B128	Header	=	0xb210
	HeaderMultihashBLAKE2B136	Header	=	0xb211
	HeaderMultihashBLAKE2B144	Header	=	0xb212
	HeaderMultihashBLAKE2B152	Header	=	0xb213
	HeaderMultihashBLAKE2B160	Header	=	0xb214
	HeaderMultihashBLAKE2B168	Header	=	0xb215
	HeaderMultihashBLAKE2B176	Header	=	0xb216
	HeaderMultihashBLAKE2B184	Header	=	0xb217
	HeaderMultihashBLAKE2B192	Header	=	0xb218
	HeaderMultihashBLAKE2B200	Header	=	0xb219
	HeaderMultihashBLAKE2B208	Header	=	0xb21a
	HeaderMultihashBLAKE2B216	Header	=	0xb21b
	HeaderMultihashBLAKE2B224	Header	=	0xb21c
	HeaderMultihashBLAKE2B232	Header	=	0xb21d
	HeaderMultihashBLAKE2B240	Header	=	0xb21e
	HeaderMultihashBLAKE2B248	Header	=	0xb21f
	HeaderMultihashBLAKE2B256	Header	=	0xb220
	HeaderMultihashBLAKE2B264	Header	=	0xb221
	HeaderMultihashBLAKE2B272	Header	=	0xb222
	HeaderMultihashBLAKE2B280	Header	=	0xb223
	HeaderMultihashBLAKE2B288	Header	=	0xb224
	HeaderMultihashBLAKE2B296	Header	=	0xb225
	HeaderMultihashBLAKE2B304	Header	=	0xb226
	HeaderMultihashBLAKE2B312	Header	=	0xb227
	HeaderMultihashBLAKE2B320	Header	=	0xb228
	HeaderMultihashBLAKE2B328	Header	=	0xb229
	HeaderMultihashBLAKE2B336	Header	=	0xb22a
	HeaderMultihashBLAKE2B344	Header	=	0xb22b
	HeaderMultihashBLAKE2B352	Header	=	0xb22c
	HeaderMultihashBLAKE2B360	Header	=	0xb22d
	HeaderMultihashBLAKE2B368	Header	=	0xb22e
	HeaderMultihashBLAKE2B376	Header	=	0xb22f
	HeaderMultihashBLAKE2B384	Header	=	0xb230
	HeaderMultihashBLAKE2B392	Header	=	0xb231
	HeaderMultihashBLAKE2B400	Header	=	0xb232
	HeaderMultihashBLAKE2B408	Header	=	0xb233
	HeaderMultihashBLAKE2B416	Header	=	0xb234
	HeaderMultihashBLAKE2B424	Header	=	0xb235
	HeaderMultihashBLAKE2B432	Header	=	0xb236
	HeaderMultihashBLAKE2B440	Header	=	0xb237
	HeaderMultihashBLAKE2B448	Header	=	0xb238
	HeaderMultihashBLAKE2B456	Header	=	0xb239
	HeaderMultihashBLAKE2B464	Header	=	0xb23a
	HeaderMultihashBLAKE2B472	Header	=	0xb23b
	HeaderMultihashBLAKE2B480	Header	=	0xb23c
	HeaderMultihashBLAKE2B488	Header	=	0xb23d
	HeaderMultihashBLAKE2B496	Header	=	0xb23e
	HeaderMultihashBLAKE2B504	Header	=	0xb23f
	HeaderMultihashBLAKE2B512	Header	=	0xb240
	HeaderMultihashBLAKE2S8	Header	=	0xb241
	HeaderMultihashBLAKE2S16	Header	=	0xb242
	HeaderMultihashBLAKE2S24	Header	=	0xb243
	HeaderMultihashBLAKE2S32	Header	=	0xb244
	HeaderMultihashBLAKE2S40	Header	=	0xb245
	HeaderMultihashBLAKE2S48	Header	=	0xb246
	HeaderMultihashBLAKE2S56	Header	=	0xb247
	HeaderMultihashBLAKE2S64	Header	=	0xb248
	HeaderMultihashBLAKE2S72	Header	=	0xb249
	HeaderMultihashBLAKE2S80	Header	=	0xb24a
	HeaderMultihashBLAKE2S88	Header	=	0xb24b
	HeaderMultihashBLAKE2S96	Header	=	0xb24c
	HeaderMultihashBLAKE2S104	Header	=	0xb24d
	HeaderMultihashBLAKE2S112	Header	=	0xb24e
	HeaderMultihashBLAKE2S120	Header	=	0xb24f
	HeaderMultihashBLAKE2S128	Header	=	0xb250
	HeaderMultihashBLAKE2S136	Header	=	0xb251
	HeaderMultihashBLAKE2S144	Header	=	0xb252
	HeaderMultihashBLAKE2S152	Header	=	0xb253
	HeaderMultihashBLAKE2S160	Header	=	0xb254
	HeaderMultihashBLAKE2S168	Header	=	0xb255
	HeaderMultihashBLAKE2S176	Header	=	0xb256
	HeaderMultihashBLAKE2S184	Header	=	0xb257
	HeaderMultihashBLAKE2S192	Header	=	0xb258
	HeaderMultihashBLAKE2S200	Header	=	0xb259
	HeaderMultihashBLAKE2S208	Header	=	0xb25a
	HeaderMultihashBLAKE2S216	Header	=	0xb25b
	HeaderMultihashBLAKE2S224	Header	=	0xb25c
	HeaderMultihashBLAKE2S232	Header	=	0xb25d
	HeaderMultihashBLAKE2S240	Header	=	0xb25e
	HeaderMultihashBLAKE2S248	Header	=	0xb25f
	HeaderMultihashBLAKE2S256	Header	=	0xb260
	HeaderMultihashSKEIN2568	Header	=	0xb301
	HeaderMultihashSKEIN25616	Header	=	0xb302
	HeaderMultihashSKEIN25624	Header	=	0xb303
	HeaderMultihashSKEIN25632	Header	=	0xb304
	HeaderMultihashSKEIN25640	Header	=	0xb305
	HeaderMultihashSKEIN25648	Header	=	0xb306
	HeaderMultihashSKEIN25656	Header	=	0xb307
	HeaderMultihashSKEIN25664	Header	=	0xb308
	HeaderMultihashSKEIN25672	Header	=	0xb309
	HeaderMultihashSKEIN25680	Header	=	0xb30a
	HeaderMultihashSKEIN25688	Header	=	0xb30b
	HeaderMultihashSKEIN25696	Header	=	0xb30c
	HeaderMultihashSKEIN256104	Header	=	0xb30d
	HeaderMultihashSKEIN256112	Header	=	0xb30e
	HeaderMultihashSKEIN256120	Header	=	0xb30f
	HeaderMultihashSKEIN256128	Header	=	0xb310
	HeaderMultihashSKEIN256136	Header	=	0xb311
	HeaderMultihashSKEIN256144	Header	=	0xb312
	HeaderMultihashSKEIN256152	Header	=	0xb313
	HeaderMultihashSKEIN256160	Header	=	0xb314
	HeaderMultihashSKEIN256168	Header	=	0xb315
	HeaderMultihashSKEIN256176	Header	=	0xb316
	HeaderMultihashSKEIN256184	Header	=	0xb317
	HeaderMultihashSKEIN256192	Header	=	0xb318
	HeaderMultihashSKEIN256200	Header	=	0xb319
	HeaderMultihashSKEIN256208	Header	=	0xb31a
	HeaderMultihashSKEIN256216	Header	=	0xb31b
	HeaderMultihashSKEIN256224	Header	=	0xb31c
	HeaderMultihashSKEIN256232	Header	=	0xb31d
	HeaderMultihashSKEIN256240	Header	=	0xb31e
	HeaderMultihashSKEIN256248	Header	=	0xb31f
	HeaderMultihashSKEIN256256	Header	=	0xb320
	HeaderMultihashSKEIN5128	Header	=	0xb321
	HeaderMultihashSKEIN51216	Header	=	0xb322
	HeaderMultihashSKEIN51224	Header	=	0xb323
	HeaderMultihashSKEIN51232	Header	=	0xb324
	HeaderMultihashSKEIN51240	Header	=	0xb325
	HeaderMultihashSKEIN51248	Header	=	0xb326
	HeaderMultihashSKEIN51256	Header	=	0xb327
	HeaderMultihashSKEIN51264	Header	=	0xb328
	HeaderMultihashSKEIN51272	Header	=	0xb329
	HeaderMultihashSKEIN51280	Header	=	0xb32a
	HeaderMultihashSKEIN51288	Header	=	0xb32b
	HeaderMultihashSKEIN51296	Header	=	0xb32c
	HeaderMultihashSKEIN512104	Header	=	0xb32d
	HeaderMultihashSKEIN512112	Header	=	0xb32e
	HeaderMultihashSKEIN512120	Header	=	0xb32f
	HeaderMultihashSKEIN512128	Header	=	0xb330
	HeaderMultihashSKEIN512136	Header	=	0xb331
	HeaderMultihashSKEIN512144	Header	=	0xb332
	HeaderMultihashSKEIN512152	Header	=	0xb333
	HeaderMultihashSKEIN512160	Header	=	0xb334
	HeaderMultihashSKEIN512168	Header	=	0xb335
	HeaderMultihashSKEIN512176	Header	=	0xb336
	HeaderMultihashSKEIN512184	Header	=	0xb337
	HeaderMultihashSKEIN512192	Header	=	0xb338
	HeaderMultihashSKEIN512200	Header	=	0xb339
	HeaderMultihashSKEIN512208	Header	=	0xb33a
	HeaderMultihashSKEIN512216	Header	=	0xb33b
	HeaderMultihashSKEIN512224	Header	=	0xb33c
	HeaderMultihashSKEIN512232	Header	=	0xb33d
	HeaderMultihashSKEIN512240	Header	=	0xb33e
	HeaderMultihashSKEIN512248	Header	=	0xb33f
	HeaderMultihashSKEIN512256	Header	=	0xb340
	HeaderMultihashSKEIN512264	Header	=	0xb341
	HeaderMultihashSKEIN512272	Header	=	0xb342
	HeaderMultihashSKEIN512280	Header	=	0xb343
	HeaderMultihashSKEIN512288	Header	=	0xb344
	HeaderMultihashSKEIN512296	Header	=	0xb345
	HeaderMultihashSKEIN512304	Header	=	0xb346
	HeaderMultihashSKEIN512312	Header	=	0xb347
	HeaderMultihashSKEIN512320	Header	=	0xb348
	HeaderMultihashSKEIN512328	Header	=	0xb349
	HeaderMultihashSKEIN512336	Header	=	0xb34a
	HeaderMultihashSKEIN512344	Header	=	0xb34b
	HeaderMultihashSKEIN512352	Header	=	0xb34c
	HeaderMultihashSKEIN512360	Header	=	0xb34d
	HeaderMultihashSKEIN512368	Header	=	0xb34e
	HeaderMultihashSKEIN512376	Header	=	0xb34f
	HeaderMultihashSKEIN512384	Header	=	0xb350
	HeaderMultihashSKEIN512392	Header	=	0xb351
	HeaderMultihashSKEIN512400	Header	=	0xb352
	HeaderMultihashSKEIN512408	Header	=	0xb353
	HeaderMultihashSKEIN512416	Header	=	0xb354
	HeaderMultihashSKEIN512424	Header	=	0xb355
	HeaderMultihashSKEIN512432	Header	=	0xb356
	HeaderMultihashSKEIN512440	Header	=	0xb357
	HeaderMultihashSKEIN512448	Header	=	0xb358
	HeaderMultihashSKEIN512456	Header	=	0xb359
	HeaderMultihashSKEIN512464	Header	=	0xb35a
	HeaderMultihashSKEIN512472	Header	=	0xb35b
	HeaderMultihashSKEIN512480	Header	=	0xb35c
	HeaderMultihashSKEIN512488	Header	=	0xb35d
	HeaderMultihashSKEIN512496	Header	=	0xb35e
	HeaderMultihashSKEIN512504	Header	=	0xb35f
	HeaderMultihashSKEIN512512	Header	=	0xb360
	HeaderMultihashSKEIN10248	Header	=	0xb361
	HeaderMultihashSKEIN102416	Header	=	0xb362
	HeaderMultihashSKEIN102424	Header	=	0xb363
	HeaderMultihashSKEIN102432	Header	=	0xb364
	HeaderMultihashSKEIN102440	Header	=	0xb365
	HeaderMultihashSKEIN102448	Header	=	0xb366
	HeaderMultihashSKEIN102456	Header	=	0xb367
	HeaderMultihashSKEIN102464	Header	=	0xb368
	HeaderMultihashSKEIN102472	Header	=	0xb369
	HeaderMultihashSKEIN102480	Header	=	0xb36a
	HeaderMultihashSKEIN102488	Header	=	0xb36b
	HeaderMultihashSKEIN102496	Header	=	0xb36c
	HeaderMultihashSKEIN1024104	Header	=	0xb36d
	HeaderMultihashSKEIN1024112	Header	=	0xb36e
	HeaderMultihashSKEIN1024120	Header	=	0xb36f
	HeaderMultihashSKEIN1024128	Header	=	0xb370
	HeaderMultihashSKEIN1024136	Header	=	0xb371
	HeaderMultihashSKEIN1024144	Header	=	0xb372
	HeaderMultihashSKEIN1024152	Header	=	0xb373
	HeaderMultihashSKEIN1024160	Header	=	0xb374
	HeaderMultihashSKEIN1024168	Header	=	0xb375
	HeaderMultihashSKEIN1024176	Header	=	0xb376
	HeaderMultihashSKEIN1024184	Header	=	0xb377
	HeaderMultihashSKEIN1024192	Header	=	0xb378
	HeaderMultihashSKEIN1024200	Header	=	0xb379
	HeaderMultihashSKEIN1024208	Header	=	0xb37a
	HeaderMultihashSKEIN1024216	Header	=	0xb37b
	HeaderMultihashSKEIN1024224	Header	=	0xb37c
	HeaderMultihashSKEIN1024232	Header	=	0xb37d
	HeaderMultihashSKEIN1024240	Header	=	0xb37e
	HeaderMultihashSKEIN1024248	Header	=	0xb37f
	HeaderMultihashSKEIN1024256	Header	=	0xb380
	HeaderMultihashSKEIN1024264	Header	=	0xb381
	HeaderMultihashSKEIN1024272	Header	=	0xb382
	HeaderMultihashSKEIN1024280	Header	=	0xb383
	HeaderMultihashSKEIN1024288	Header	=	0xb384
	HeaderMultihashSKEIN1024296	Header	=	0xb385
	HeaderMultihashSKEIN1024304	Header	=	0xb386
	HeaderMultihashSKEIN1024312	Header	=	0xb387
	HeaderMultihashSKEIN1024320	Header	=	0xb388
	HeaderMultihashSKEIN1024328	Header	=	0xb389
	HeaderMultihashSKEIN1024336	Header	=	0xb38a
	HeaderMultihashSKEIN1024344	Header	=	0xb38b
	HeaderMultihashSKEIN1024352	Header	=	0xb38c
	HeaderMultihashSKEIN1024360	Header	=	0xb38d
	HeaderMultihashSKEIN1024368	Header	=	0xb38e
	HeaderMultihashSKEIN1024376	Header	=	0xb38f
	HeaderMultihashSKEIN1024384	Header	=	0xb390
	HeaderMultihashSKEIN1024392	Header	=	0xb391
	HeaderMultihashSKEIN1024400	Header	=	0xb392
	HeaderMultihashSKEIN1024408	Header	=	0xb393
	HeaderMultihashSKEIN1024416	Header	=	0xb394
	HeaderMultihashSKEIN1024424	Header	=	0xb395
	HeaderMultihashSKEIN1024432	Header	=	0xb396
	HeaderMultihashSKEIN1024440	Header	=	0xb397
	HeaderMultihashSKEIN1024448	Header	=	0xb398
	HeaderMultihashSKEIN1024456	Header	=	0xb399
	HeaderMultihashSKEIN1024464	Header	=	0xb39a
	HeaderMultihashSKEIN1024472	Header	=	0xb39b
	HeaderMultihashSKEIN1024480	Header	=	0xb39c
	HeaderMultihashSKEIN1024488	Header	=	0xb39d
	HeaderMultihashSKEIN1024496	Header	=	0xb39e
	HeaderMultihashSKEIN1024504	Header	=	0xb39f
	HeaderMultihashSKEIN1024512	Header	=	0xb3a0
	HeaderMultihashSKEIN1024520	Header	=	0xb3a1
	HeaderMultihashSKEIN1024528	Header	=	0xb3a2
	HeaderMultihashSKEIN1024536	Header	=	0xb3a3
	HeaderMultihashSKEIN1024544	Header	=	0xb3a4
	HeaderMultihashSKEIN1024552	Header	=	0xb3a5
	HeaderMultihashSKEIN1024560	Header	=	0xb3a6
	HeaderMultihashSKEIN1024568	Header	=	0xb3a7
	HeaderMultihashSKEIN1024576	Header	=	0xb3a8
	HeaderMultihashSKEIN1024584	Header	=	0xb3a9
	HeaderMultihashSKEIN1024592	Header	=	0xb3aa
	HeaderMultihashSKEIN1024600	Header	=	0xb3ab
	HeaderMultihashSKEIN1024608	Header	=	0xb3ac
	HeaderMultihashSKEIN1024616	Header	=	0xb3ad
	HeaderMultihashSKEIN1024624	Header	=	0xb3ae
	HeaderMultihashSKEIN1024632	Header	=	0xb3af
	HeaderMultihashSKEIN1024640	Header	=	0xb3b0
	HeaderMultihashSKEIN1024648	Header	=	0xb3b1
	HeaderMultihashSKEIN1024656	Header	=	0xb3b2
	HeaderMultihashSKEIN1024664	Header	=	0xb3b3
	HeaderMultihashSKEIN1024672	Header	=	0xb3b4
	HeaderMultihashSKEIN1024680	Header	=	0xb3b5
	HeaderMultihashSKEIN1024688	Header	=	0xb3b6
	HeaderMultihashSKEIN1024696	Header	=	0xb3b7
	HeaderMultihashSKEIN1024704	Header	=	0xb3b8
	HeaderMultihashSKEIN1024712	Header	=	0xb3b9
	HeaderMultihashSKEIN1024720	Header	=	0xb3ba
	HeaderMultihashSKEIN1024728	Header	=	0xb3bb
	HeaderMultihashSKEIN1024736	Header	=	0xb3bc
	HeaderMultihashSKEIN1024744	Header	=	0xb3bd
	HeaderMultihashSKEIN1024752	Header	=	0xb3be
	HeaderMultihashSKEIN1024760	Header	=	0xb3bf
	HeaderMultihashSKEIN1024768	Header	=	0xb3c0
	HeaderMultihashSKEIN1024776	Header	=	0xb3c1
	HeaderMultihashSKEIN1024784	Header	=	0xb3c2
	HeaderMultihashSKEIN1024792	Header	=	0xb3c3
	HeaderMultihashSKEIN1024800	Header	=	0xb3c4
	HeaderMultihashSKEIN1024808	Header	=	0xb3c5
	HeaderMultihashSKEIN1024816	Header	=	0xb3c6
	HeaderMultihashSKEIN1024824	Header	=	0xb3c7
	HeaderMultihashSKEIN1024832	Header	=	0xb3c8
	HeaderMultihashSKEIN1024840	Header	=	0xb3c9
	HeaderMultihashSKEIN1024848	Header	=	0xb3ca
	HeaderMultihashSKEIN1024856	Header	=	0xb3cb
	HeaderMultihashSKEIN1024864	Header	=	0xb3cc
	HeaderMultihashSKEIN1024872	Header	=	0xb3cd
	HeaderMultihashSKEIN1024880	Header	=	0xb3ce
	HeaderMultihashSKEIN1024888	Header	=	0xb3cf
	HeaderMultihashSKEIN1024896	Header	=	0xb3d0
	HeaderMultihashSKEIN1024904	Header	=	0xb3d1
	HeaderMultihashSKEIN1024912	Header	=	0xb3d2
	HeaderMultihashSKEIN1024920	Header	=	0xb3d3
	HeaderMultihashSKEIN1024928	Header	=	0xb3d4
	HeaderMultihashSKEIN1024936	Header	=	0xb3d5
	HeaderMultihashSKEIN1024944	Header	=	0xb3d6
	HeaderMultihashSKEIN1024952	Header	=	0xb3d7
	HeaderMultihashSKEIN1024960	Header	=	0xb3d8
	HeaderMultihashSKEIN1024968	Header	=	0xb3d9
	HeaderMultihashSKEIN1024976	Header	=	0xb3da
	HeaderMultihashSKEIN1024984	Header	=	0xb3db
	HeaderMultihashSKEIN1024992	Header	=	0xb3dc
	HeaderMultihashSKEIN10241000	Header	=	0xb3dd
	HeaderMultihashSKEIN10241008	Header	=	0xb3de
	HeaderMultihashSKEIN10241016	Header	=	0xb3df
	HeaderMultihashSKEIN10241024	Header	=	0xb3e0
	HeaderHashXXH32	Header	=	0xb3e1
	HeaderHashXXH64	Header	=	0xb3e2
	HeaderHashXXH364	Header	=	0xb3e3
	HeaderHashXXH3128	Header	=	0xb3e4
	HeaderMultihashPOSEIDONBLS12381A2FC1	Header	=	0xb401
	HeaderMultihashPOSEIDONBLS12381A2FC1SC	Header	=	0xb402
	HeaderIpldRDFC1	Header	=	0xb403
	HeaderSerializationSSZ	Header	=	0xb501
	HeaderMultihashSSZSHA2256BMT	Header	=	0xb502
	HeaderIpldJSONJCS	Header	=	0xb601
	HeaderSofthashISCC	Header	=	0xcc01
	HeaderZeroxcertZEROXCERTIMPRINT256	Header	=	0xce11
	HeaderVarsigNONSTANDARDSIG	Header	=	0xd000
	HeaderVarsigES256K	Header	=	0xd0e7
	HeaderVarsigBLS12381G1SIG	Header	=	0xd0ea
	HeaderVarsigBLS12381G2SIG	Header	=	0xd0eb
	HeaderVarsigEDDSA	Header	=	0xd0ed
	HeaderVarsigEIP191	Header	=	0xd191
	HeaderKeyJWKJCSPUB	Header	=	0xeb51
	HeaderFilecoinFILCOMMITMENTUNSEALED	Header	=	0xf101
	HeaderFilecoinFILCOMMITMENTSEALED	Header	=	0xf102
	HeaderMultiaddrPLAINTEXTV2	Header	=	0x706c61
	HeaderHolochainHOLOCHAINADRV0	Header	=	0x807124
	HeaderHolochainHOLOCHAINADRV1	Header	=	0x817124
	HeaderHolochainHOLOCHAINKEYV0	Header	=	0x947124
	HeaderHolochainHOLOCHAINKEYV1	Header	=	0x957124
	HeaderHolochainHOLOCHAINSIGV0	Header	=	0xa27124
	HeaderHolochainHOLOCHAINSIGV1	Header	=	0xa37124
	HeaderNamespaceSKYNETNS	Header	=	0xb19910
	HeaderNamespaceARWEAVENS	Header	=	0xb29910
	HeaderNamespaceSUBSPACENS	Header	=	0xb39910
	HeaderNamespaceKUMANDRANS	Header	=	0xb49910
	HeaderVarsigES256	Header	=	0xd01200
	HeaderVarsigES284	Header	=	0xd01201
	HeaderVarsigES512	Header	=	0xd01202
	HeaderVarsigRS256	Header	=	0xd01205
	HeaderMultiaddrSCION	Header	=	0xd02000


	IdentifierUNKNOWN			Identifier	= ""
	IdentifierMultihashIDENTITY	Identifier	=	"MultihashIDENTITY"
	IdentifierCidCIDV1	Identifier	=	"CidCIDV1"
	IdentifierCidCIDV2	Identifier	=	"CidCIDV2"
	IdentifierCidCIDV3	Identifier	=	"CidCIDV3"
	IdentifierMultiaddrIP4	Identifier	=	"MultiaddrIP4"
	IdentifierMultiaddrTCP	Identifier	=	"MultiaddrTCP"
	IdentifierMultihashSHA1	Identifier	=	"MultihashSHA1"
	IdentifierMultihashSHA2256	Identifier	=	"MultihashSHA2256"
	IdentifierMultihashSHA2512	Identifier	=	"MultihashSHA2512"
	IdentifierMultihashSHA3512	Identifier	=	"MultihashSHA3512"
	IdentifierMultihashSHA3384	Identifier	=	"MultihashSHA3384"
	IdentifierMultihashSHA3256	Identifier	=	"MultihashSHA3256"
	IdentifierMultihashSHA3224	Identifier	=	"MultihashSHA3224"
	IdentifierMultihashSHAKE128	Identifier	=	"MultihashSHAKE128"
	IdentifierMultihashSHAKE256	Identifier	=	"MultihashSHAKE256"
	IdentifierMultihashKECCAK224	Identifier	=	"MultihashKECCAK224"
	IdentifierMultihashKECCAK256	Identifier	=	"MultihashKECCAK256"
	IdentifierMultihashKECCAK384	Identifier	=	"MultihashKECCAK384"
	IdentifierMultihashKECCAK512	Identifier	=	"MultihashKECCAK512"
	IdentifierMultihashBLAKE3	Identifier	=	"MultihashBLAKE3"
	IdentifierMultihashSHA2384	Identifier	=	"MultihashSHA2384"
	IdentifierMultiaddrDCCP	Identifier	=	"MultiaddrDCCP"
	IdentifierHashMURMUR3X6464	Identifier	=	"HashMURMUR3X6464"
	IdentifierHashMURMUR332	Identifier	=	"HashMURMUR332"
	IdentifierMultiaddrIP6	Identifier	=	"MultiaddrIP6"
	IdentifierMultiaddrIP6ZONE	Identifier	=	"MultiaddrIP6ZONE"
	IdentifierMultiaddrIPCIDR	Identifier	=	"MultiaddrIPCIDR"
	IdentifierNamespacePATH	Identifier	=	"NamespacePATH"
	IdentifierMultiformatMULTICODEC	Identifier	=	"MultiformatMULTICODEC"
	IdentifierMultiformatMULTIHASH	Identifier	=	"MultiformatMULTIHASH"
	IdentifierMultiformatMULTIADDR	Identifier	=	"MultiformatMULTIADDR"
	IdentifierMultiformatMULTIBASE	Identifier	=	"MultiformatMULTIBASE"
	IdentifierMultiformatVARSIG	Identifier	=	"MultiformatVARSIG"
	IdentifierMultiaddrDNS	Identifier	=	"MultiaddrDNS"
	IdentifierMultiaddrDNS4	Identifier	=	"MultiaddrDNS4"
	IdentifierMultiaddrDNS6	Identifier	=	"MultiaddrDNS6"
	IdentifierMultiaddrDNSADDR	Identifier	=	"MultiaddrDNSADDR"
	IdentifierSerializationPROTOBUF	Identifier	=	"SerializationPROTOBUF"
	IdentifierIpldCBOR	Identifier	=	"IpldCBOR"
	IdentifierIpldRAW	Identifier	=	"IpldRAW"
	IdentifierMultihashDBLSHA2256	Identifier	=	"MultihashDBLSHA2256"
	IdentifierSerializationRLP	Identifier	=	"SerializationRLP"
	IdentifierSerializationBENCODE	Identifier	=	"SerializationBENCODE"
	IdentifierIpldDAGPB	Identifier	=	"IpldDAGPB"
	IdentifierIpldDAGCBOR	Identifier	=	"IpldDAGCBOR"
	IdentifierIpldLIBP2PKEY	Identifier	=	"IpldLIBP2PKEY"
	IdentifierIpldGITRAW	Identifier	=	"IpldGITRAW"
	IdentifierIpldTORRENTINFO	Identifier	=	"IpldTORRENTINFO"
	IdentifierIpldTORRENTFILE	Identifier	=	"IpldTORRENTFILE"
	IdentifierIpldLEOFCOINBLOCK	Identifier	=	"IpldLEOFCOINBLOCK"
	IdentifierIpldLEOFCOINTX	Identifier	=	"IpldLEOFCOINTX"
	IdentifierIpldLEOFCOINPR	Identifier	=	"IpldLEOFCOINPR"
	IdentifierMultiaddrSCTP	Identifier	=	"MultiaddrSCTP"
	IdentifierIpldDAGJOSE	Identifier	=	"IpldDAGJOSE"
	IdentifierIpldDAGCOSE	Identifier	=	"IpldDAGCOSE"
	IdentifierNamespaceLBRY	Identifier	=	"NamespaceLBRY"
	IdentifierIpldETHBLOCK	Identifier	=	"IpldETHBLOCK"
	IdentifierIpldETHBLOCKLIST	Identifier	=	"IpldETHBLOCKLIST"
	IdentifierIpldETHTXTRIE	Identifier	=	"IpldETHTXTRIE"
	IdentifierIpldETHTX	Identifier	=	"IpldETHTX"
	IdentifierIpldETHTXRECEIPTTRIE	Identifier	=	"IpldETHTXRECEIPTTRIE"
	IdentifierIpldETHTXRECEIPT	Identifier	=	"IpldETHTXRECEIPT"
	IdentifierIpldETHSTATETRIE	Identifier	=	"IpldETHSTATETRIE"
	IdentifierIpldETHACCOUNTSNAPSHOT	Identifier	=	"IpldETHACCOUNTSNAPSHOT"
	IdentifierIpldETHSTORAGETRIE	Identifier	=	"IpldETHSTORAGETRIE"
	IdentifierIpldETHRECEIPTLOGTRIE	Identifier	=	"IpldETHRECEIPTLOGTRIE"
	IdentifierIpldETHRECEIPTLOG	Identifier	=	"IpldETHRECEIPTLOG"
	IdentifierKeyAES128	Identifier	=	"KeyAES128"
	IdentifierKeyAES192	Identifier	=	"KeyAES192"
	IdentifierKeyAES256	Identifier	=	"KeyAES256"
	IdentifierKeyCHACHA128	Identifier	=	"KeyCHACHA128"
	IdentifierKeyCHACHA256	Identifier	=	"KeyCHACHA256"
	IdentifierIpldBITCOINBLOCK	Identifier	=	"IpldBITCOINBLOCK"
	IdentifierIpldBITCOINTX	Identifier	=	"IpldBITCOINTX"
	IdentifierIpldBITCOINWITNESSCOMMITMENT	Identifier	=	"IpldBITCOINWITNESSCOMMITMENT"
	IdentifierIpldZCASHBLOCK	Identifier	=	"IpldZCASHBLOCK"
	IdentifierIpldZCASHTX	Identifier	=	"IpldZCASHTX"
	IdentifierMultiformatCAIP50	Identifier	=	"MultiformatCAIP50"
	IdentifierNamespaceSTREAMID	Identifier	=	"NamespaceSTREAMID"
	IdentifierIpldSTELLARBLOCK	Identifier	=	"IpldSTELLARBLOCK"
	IdentifierIpldSTELLARTX	Identifier	=	"IpldSTELLARTX"
	IdentifierMultihashMD4	Identifier	=	"MultihashMD4"
	IdentifierMultihashMD5	Identifier	=	"MultihashMD5"
	IdentifierIpldDECREDBLOCK	Identifier	=	"IpldDECREDBLOCK"
	IdentifierIpldDECREDTX	Identifier	=	"IpldDECREDTX"
	IdentifierNamespaceIPLD	Identifier	=	"NamespaceIPLD"
	IdentifierNamespaceIPFS	Identifier	=	"NamespaceIPFS"
	IdentifierNamespaceSWARM	Identifier	=	"NamespaceSWARM"
	IdentifierNamespaceIPNS	Identifier	=	"NamespaceIPNS"
	IdentifierNamespaceZERONET	Identifier	=	"NamespaceZERONET"
	IdentifierKeySECP256K1PUB	Identifier	=	"KeySECP256K1PUB"
	IdentifierNamespaceDNSLINK	Identifier	=	"NamespaceDNSLINK"
	IdentifierKeyBLS12381G1PUB	Identifier	=	"KeyBLS12381G1PUB"
	IdentifierKeyBLS12381G2PUB	Identifier	=	"KeyBLS12381G2PUB"
	IdentifierKeyX25519PUB	Identifier	=	"KeyX25519PUB"
	IdentifierKeyED25519PUB	Identifier	=	"KeyED25519PUB"
	IdentifierKeyBLS12381G1G2PUB	Identifier	=	"KeyBLS12381G1G2PUB"
	IdentifierKeySR25519PUB	Identifier	=	"KeySR25519PUB"
	IdentifierIpldDASHBLOCK	Identifier	=	"IpldDASHBLOCK"
	IdentifierIpldDASHTX	Identifier	=	"IpldDASHTX"
	IdentifierIpldSWARMMANIFEST	Identifier	=	"IpldSWARMMANIFEST"
	IdentifierIpldSWARMFEED	Identifier	=	"IpldSWARMFEED"
	IdentifierIpldBEESON	Identifier	=	"IpldBEESON"
	IdentifierMultiaddrUDP	Identifier	=	"MultiaddrUDP"
	IdentifierMultiaddrP2PWEBRTCSTAR	Identifier	=	"MultiaddrP2PWEBRTCSTAR"
	IdentifierMultiaddrP2PWEBRTCDIRECT	Identifier	=	"MultiaddrP2PWEBRTCDIRECT"
	IdentifierMultiaddrP2PSTARDUST	Identifier	=	"MultiaddrP2PSTARDUST"
	IdentifierMultiaddrWEBRTCDIRECT	Identifier	=	"MultiaddrWEBRTCDIRECT"
	IdentifierMultiaddrWEBRTC	Identifier	=	"MultiaddrWEBRTC"
	IdentifierMultiaddrP2PCIRCUIT	Identifier	=	"MultiaddrP2PCIRCUIT"
	IdentifierIpldDAGJSON	Identifier	=	"IpldDAGJSON"
	IdentifierMultiaddrUDT	Identifier	=	"MultiaddrUDT"
	IdentifierMultiaddrUTP	Identifier	=	"MultiaddrUTP"
	IdentifierHashCRC32	Identifier	=	"HashCRC32"
	IdentifierHashCRC64ECMA	Identifier	=	"HashCRC64ECMA"
	IdentifierMultiaddrUNIX	Identifier	=	"MultiaddrUNIX"
	IdentifierMultiaddrTHREAD	Identifier	=	"MultiaddrTHREAD"
	IdentifierMultiaddrP2P	Identifier	=	"MultiaddrP2P"
	IdentifierMultiaddrHTTPS	Identifier	=	"MultiaddrHTTPS"
	IdentifierMultiaddrONION	Identifier	=	"MultiaddrONION"
	IdentifierMultiaddrONION3	Identifier	=	"MultiaddrONION3"
	IdentifierMultiaddrGARLIC64	Identifier	=	"MultiaddrGARLIC64"
	IdentifierMultiaddrGARLIC32	Identifier	=	"MultiaddrGARLIC32"
	IdentifierMultiaddrTLS	Identifier	=	"MultiaddrTLS"
	IdentifierMultiaddrSNI	Identifier	=	"MultiaddrSNI"
	IdentifierMultiaddrNOISE	Identifier	=	"MultiaddrNOISE"
	IdentifierMultiaddrSHS	Identifier	=	"MultiaddrSHS"
	IdentifierMultiaddrQUIC	Identifier	=	"MultiaddrQUIC"
	IdentifierMultiaddrQUICV1	Identifier	=	"MultiaddrQUICV1"
	IdentifierMultiaddrWEBTRANSPORT	Identifier	=	"MultiaddrWEBTRANSPORT"
	IdentifierMultiaddrCERTHASH	Identifier	=	"MultiaddrCERTHASH"
	IdentifierMultiaddrWS	Identifier	=	"MultiaddrWS"
	IdentifierMultiaddrWSS	Identifier	=	"MultiaddrWSS"
	IdentifierMultiaddrP2PWEBSOCKETSTAR	Identifier	=	"MultiaddrP2PWEBSOCKETSTAR"
	IdentifierMultiaddrHTTP	Identifier	=	"MultiaddrHTTP"
	IdentifierIpldSWHID1SNP	Identifier	=	"IpldSWHID1SNP"
	IdentifierIpldJSON	Identifier	=	"IpldJSON"
	IdentifierSerializationMESSAGEPACK	Identifier	=	"SerializationMESSAGEPACK"
	IdentifierSerializationCAR	Identifier	=	"SerializationCAR"
	IdentifierSerializationIPNSRECORD	Identifier	=	"SerializationIPNSRECORD"
	IdentifierLibp2pLIBP2PPEERRECORD	Identifier	=	"Libp2pLIBP2PPEERRECORD"
	IdentifierLibp2pLIBP2PRELAYRSVP	Identifier	=	"Libp2pLIBP2PRELAYRSVP"
	IdentifierLibp2pMEMORYTRANSPORT	Identifier	=	"Libp2pMEMORYTRANSPORT"
	IdentifierSerializationCARINDEXSORTED	Identifier	=	"SerializationCARINDEXSORTED"
	IdentifierSerializationCARMULTIHASHINDEXSORTED	Identifier	=	"SerializationCARMULTIHASHINDEXSORTED"
	IdentifierTransportTRANSPORTBITSWAP	Identifier	=	"TransportTRANSPORTBITSWAP"
	IdentifierTransportTRANSPORTGRAPHSYNCFILECOINV1	Identifier	=	"TransportTRANSPORTGRAPHSYNCFILECOINV1"
	IdentifierTransportTRANSPORTIPFSGATEWAYHTTP	Identifier	=	"TransportTRANSPORTIPFSGATEWAYHTTP"
	IdentifierMultiformatMULTIDID	Identifier	=	"MultiformatMULTIDID"
	IdentifierMultihashSHA2256TRUNC254PADDED	Identifier	=	"MultihashSHA2256TRUNC254PADDED"
	IdentifierMultihashSHA2224	Identifier	=	"MultihashSHA2224"
	IdentifierMultihashSHA2512224	Identifier	=	"MultihashSHA2512224"
	IdentifierMultihashSHA2512256	Identifier	=	"MultihashSHA2512256"
	IdentifierHashMURMUR3X64128	Identifier	=	"HashMURMUR3X64128"
	IdentifierMultihashRIPEMD128	Identifier	=	"MultihashRIPEMD128"
	IdentifierMultihashRIPEMD160	Identifier	=	"MultihashRIPEMD160"
	IdentifierMultihashRIPEMD256	Identifier	=	"MultihashRIPEMD256"
	IdentifierMultihashRIPEMD320	Identifier	=	"MultihashRIPEMD320"
	IdentifierMultihashX11	Identifier	=	"MultihashX11"
	IdentifierKeyP256PUB	Identifier	=	"KeyP256PUB"
	IdentifierKeyP384PUB	Identifier	=	"KeyP384PUB"
	IdentifierKeyP521PUB	Identifier	=	"KeyP521PUB"
	IdentifierKeyED448PUB	Identifier	=	"KeyED448PUB"
	IdentifierKeyX448PUB	Identifier	=	"KeyX448PUB"
	IdentifierKeyRSAPUB	Identifier	=	"KeyRSAPUB"
	IdentifierKeySM2PUB	Identifier	=	"KeySM2PUB"
	IdentifierKeyED25519PRIV	Identifier	=	"KeyED25519PRIV"
	IdentifierKeySECP256K1PRIV	Identifier	=	"KeySECP256K1PRIV"
	IdentifierKeyX25519PRIV	Identifier	=	"KeyX25519PRIV"
	IdentifierKeySR25519PRIV	Identifier	=	"KeySR25519PRIV"
	IdentifierKeyRSAPRIV	Identifier	=	"KeyRSAPRIV"
	IdentifierKeyP256PRIV	Identifier	=	"KeyP256PRIV"
	IdentifierKeyP384PRIV	Identifier	=	"KeyP384PRIV"
	IdentifierKeyP521PRIV	Identifier	=	"KeyP521PRIV"
	IdentifierMultihashKANGAROOTWELVE	Identifier	=	"MultihashKANGAROOTWELVE"
	IdentifierEncryptionAESGCM256	Identifier	=	"EncryptionAESGCM256"
	IdentifierMultiaddrSILVERPINE	Identifier	=	"MultiaddrSILVERPINE"
	IdentifierMultihashSM3256	Identifier	=	"MultihashSM3256"
	IdentifierHashSHA256A	Identifier	=	"HashSHA256A"
	IdentifierMultihashBLAKE2B8	Identifier	=	"MultihashBLAKE2B8"
	IdentifierMultihashBLAKE2B16	Identifier	=	"MultihashBLAKE2B16"
	IdentifierMultihashBLAKE2B24	Identifier	=	"MultihashBLAKE2B24"
	IdentifierMultihashBLAKE2B32	Identifier	=	"MultihashBLAKE2B32"
	IdentifierMultihashBLAKE2B40	Identifier	=	"MultihashBLAKE2B40"
	IdentifierMultihashBLAKE2B48	Identifier	=	"MultihashBLAKE2B48"
	IdentifierMultihashBLAKE2B56	Identifier	=	"MultihashBLAKE2B56"
	IdentifierMultihashBLAKE2B64	Identifier	=	"MultihashBLAKE2B64"
	IdentifierMultihashBLAKE2B72	Identifier	=	"MultihashBLAKE2B72"
	IdentifierMultihashBLAKE2B80	Identifier	=	"MultihashBLAKE2B80"
	IdentifierMultihashBLAKE2B88	Identifier	=	"MultihashBLAKE2B88"
	IdentifierMultihashBLAKE2B96	Identifier	=	"MultihashBLAKE2B96"
	IdentifierMultihashBLAKE2B104	Identifier	=	"MultihashBLAKE2B104"
	IdentifierMultihashBLAKE2B112	Identifier	=	"MultihashBLAKE2B112"
	IdentifierMultihashBLAKE2B120	Identifier	=	"MultihashBLAKE2B120"
	IdentifierMultihashBLAKE2B128	Identifier	=	"MultihashBLAKE2B128"
	IdentifierMultihashBLAKE2B136	Identifier	=	"MultihashBLAKE2B136"
	IdentifierMultihashBLAKE2B144	Identifier	=	"MultihashBLAKE2B144"
	IdentifierMultihashBLAKE2B152	Identifier	=	"MultihashBLAKE2B152"
	IdentifierMultihashBLAKE2B160	Identifier	=	"MultihashBLAKE2B160"
	IdentifierMultihashBLAKE2B168	Identifier	=	"MultihashBLAKE2B168"
	IdentifierMultihashBLAKE2B176	Identifier	=	"MultihashBLAKE2B176"
	IdentifierMultihashBLAKE2B184	Identifier	=	"MultihashBLAKE2B184"
	IdentifierMultihashBLAKE2B192	Identifier	=	"MultihashBLAKE2B192"
	IdentifierMultihashBLAKE2B200	Identifier	=	"MultihashBLAKE2B200"
	IdentifierMultihashBLAKE2B208	Identifier	=	"MultihashBLAKE2B208"
	IdentifierMultihashBLAKE2B216	Identifier	=	"MultihashBLAKE2B216"
	IdentifierMultihashBLAKE2B224	Identifier	=	"MultihashBLAKE2B224"
	IdentifierMultihashBLAKE2B232	Identifier	=	"MultihashBLAKE2B232"
	IdentifierMultihashBLAKE2B240	Identifier	=	"MultihashBLAKE2B240"
	IdentifierMultihashBLAKE2B248	Identifier	=	"MultihashBLAKE2B248"
	IdentifierMultihashBLAKE2B256	Identifier	=	"MultihashBLAKE2B256"
	IdentifierMultihashBLAKE2B264	Identifier	=	"MultihashBLAKE2B264"
	IdentifierMultihashBLAKE2B272	Identifier	=	"MultihashBLAKE2B272"
	IdentifierMultihashBLAKE2B280	Identifier	=	"MultihashBLAKE2B280"
	IdentifierMultihashBLAKE2B288	Identifier	=	"MultihashBLAKE2B288"
	IdentifierMultihashBLAKE2B296	Identifier	=	"MultihashBLAKE2B296"
	IdentifierMultihashBLAKE2B304	Identifier	=	"MultihashBLAKE2B304"
	IdentifierMultihashBLAKE2B312	Identifier	=	"MultihashBLAKE2B312"
	IdentifierMultihashBLAKE2B320	Identifier	=	"MultihashBLAKE2B320"
	IdentifierMultihashBLAKE2B328	Identifier	=	"MultihashBLAKE2B328"
	IdentifierMultihashBLAKE2B336	Identifier	=	"MultihashBLAKE2B336"
	IdentifierMultihashBLAKE2B344	Identifier	=	"MultihashBLAKE2B344"
	IdentifierMultihashBLAKE2B352	Identifier	=	"MultihashBLAKE2B352"
	IdentifierMultihashBLAKE2B360	Identifier	=	"MultihashBLAKE2B360"
	IdentifierMultihashBLAKE2B368	Identifier	=	"MultihashBLAKE2B368"
	IdentifierMultihashBLAKE2B376	Identifier	=	"MultihashBLAKE2B376"
	IdentifierMultihashBLAKE2B384	Identifier	=	"MultihashBLAKE2B384"
	IdentifierMultihashBLAKE2B392	Identifier	=	"MultihashBLAKE2B392"
	IdentifierMultihashBLAKE2B400	Identifier	=	"MultihashBLAKE2B400"
	IdentifierMultihashBLAKE2B408	Identifier	=	"MultihashBLAKE2B408"
	IdentifierMultihashBLAKE2B416	Identifier	=	"MultihashBLAKE2B416"
	IdentifierMultihashBLAKE2B424	Identifier	=	"MultihashBLAKE2B424"
	IdentifierMultihashBLAKE2B432	Identifier	=	"MultihashBLAKE2B432"
	IdentifierMultihashBLAKE2B440	Identifier	=	"MultihashBLAKE2B440"
	IdentifierMultihashBLAKE2B448	Identifier	=	"MultihashBLAKE2B448"
	IdentifierMultihashBLAKE2B456	Identifier	=	"MultihashBLAKE2B456"
	IdentifierMultihashBLAKE2B464	Identifier	=	"MultihashBLAKE2B464"
	IdentifierMultihashBLAKE2B472	Identifier	=	"MultihashBLAKE2B472"
	IdentifierMultihashBLAKE2B480	Identifier	=	"MultihashBLAKE2B480"
	IdentifierMultihashBLAKE2B488	Identifier	=	"MultihashBLAKE2B488"
	IdentifierMultihashBLAKE2B496	Identifier	=	"MultihashBLAKE2B496"
	IdentifierMultihashBLAKE2B504	Identifier	=	"MultihashBLAKE2B504"
	IdentifierMultihashBLAKE2B512	Identifier	=	"MultihashBLAKE2B512"
	IdentifierMultihashBLAKE2S8	Identifier	=	"MultihashBLAKE2S8"
	IdentifierMultihashBLAKE2S16	Identifier	=	"MultihashBLAKE2S16"
	IdentifierMultihashBLAKE2S24	Identifier	=	"MultihashBLAKE2S24"
	IdentifierMultihashBLAKE2S32	Identifier	=	"MultihashBLAKE2S32"
	IdentifierMultihashBLAKE2S40	Identifier	=	"MultihashBLAKE2S40"
	IdentifierMultihashBLAKE2S48	Identifier	=	"MultihashBLAKE2S48"
	IdentifierMultihashBLAKE2S56	Identifier	=	"MultihashBLAKE2S56"
	IdentifierMultihashBLAKE2S64	Identifier	=	"MultihashBLAKE2S64"
	IdentifierMultihashBLAKE2S72	Identifier	=	"MultihashBLAKE2S72"
	IdentifierMultihashBLAKE2S80	Identifier	=	"MultihashBLAKE2S80"
	IdentifierMultihashBLAKE2S88	Identifier	=	"MultihashBLAKE2S88"
	IdentifierMultihashBLAKE2S96	Identifier	=	"MultihashBLAKE2S96"
	IdentifierMultihashBLAKE2S104	Identifier	=	"MultihashBLAKE2S104"
	IdentifierMultihashBLAKE2S112	Identifier	=	"MultihashBLAKE2S112"
	IdentifierMultihashBLAKE2S120	Identifier	=	"MultihashBLAKE2S120"
	IdentifierMultihashBLAKE2S128	Identifier	=	"MultihashBLAKE2S128"
	IdentifierMultihashBLAKE2S136	Identifier	=	"MultihashBLAKE2S136"
	IdentifierMultihashBLAKE2S144	Identifier	=	"MultihashBLAKE2S144"
	IdentifierMultihashBLAKE2S152	Identifier	=	"MultihashBLAKE2S152"
	IdentifierMultihashBLAKE2S160	Identifier	=	"MultihashBLAKE2S160"
	IdentifierMultihashBLAKE2S168	Identifier	=	"MultihashBLAKE2S168"
	IdentifierMultihashBLAKE2S176	Identifier	=	"MultihashBLAKE2S176"
	IdentifierMultihashBLAKE2S184	Identifier	=	"MultihashBLAKE2S184"
	IdentifierMultihashBLAKE2S192	Identifier	=	"MultihashBLAKE2S192"
	IdentifierMultihashBLAKE2S200	Identifier	=	"MultihashBLAKE2S200"
	IdentifierMultihashBLAKE2S208	Identifier	=	"MultihashBLAKE2S208"
	IdentifierMultihashBLAKE2S216	Identifier	=	"MultihashBLAKE2S216"
	IdentifierMultihashBLAKE2S224	Identifier	=	"MultihashBLAKE2S224"
	IdentifierMultihashBLAKE2S232	Identifier	=	"MultihashBLAKE2S232"
	IdentifierMultihashBLAKE2S240	Identifier	=	"MultihashBLAKE2S240"
	IdentifierMultihashBLAKE2S248	Identifier	=	"MultihashBLAKE2S248"
	IdentifierMultihashBLAKE2S256	Identifier	=	"MultihashBLAKE2S256"
	IdentifierMultihashSKEIN2568	Identifier	=	"MultihashSKEIN2568"
	IdentifierMultihashSKEIN25616	Identifier	=	"MultihashSKEIN25616"
	IdentifierMultihashSKEIN25624	Identifier	=	"MultihashSKEIN25624"
	IdentifierMultihashSKEIN25632	Identifier	=	"MultihashSKEIN25632"
	IdentifierMultihashSKEIN25640	Identifier	=	"MultihashSKEIN25640"
	IdentifierMultihashSKEIN25648	Identifier	=	"MultihashSKEIN25648"
	IdentifierMultihashSKEIN25656	Identifier	=	"MultihashSKEIN25656"
	IdentifierMultihashSKEIN25664	Identifier	=	"MultihashSKEIN25664"
	IdentifierMultihashSKEIN25672	Identifier	=	"MultihashSKEIN25672"
	IdentifierMultihashSKEIN25680	Identifier	=	"MultihashSKEIN25680"
	IdentifierMultihashSKEIN25688	Identifier	=	"MultihashSKEIN25688"
	IdentifierMultihashSKEIN25696	Identifier	=	"MultihashSKEIN25696"
	IdentifierMultihashSKEIN256104	Identifier	=	"MultihashSKEIN256104"
	IdentifierMultihashSKEIN256112	Identifier	=	"MultihashSKEIN256112"
	IdentifierMultihashSKEIN256120	Identifier	=	"MultihashSKEIN256120"
	IdentifierMultihashSKEIN256128	Identifier	=	"MultihashSKEIN256128"
	IdentifierMultihashSKEIN256136	Identifier	=	"MultihashSKEIN256136"
	IdentifierMultihashSKEIN256144	Identifier	=	"MultihashSKEIN256144"
	IdentifierMultihashSKEIN256152	Identifier	=	"MultihashSKEIN256152"
	IdentifierMultihashSKEIN256160	Identifier	=	"MultihashSKEIN256160"
	IdentifierMultihashSKEIN256168	Identifier	=	"MultihashSKEIN256168"
	IdentifierMultihashSKEIN256176	Identifier	=	"MultihashSKEIN256176"
	IdentifierMultihashSKEIN256184	Identifier	=	"MultihashSKEIN256184"
	IdentifierMultihashSKEIN256192	Identifier	=	"MultihashSKEIN256192"
	IdentifierMultihashSKEIN256200	Identifier	=	"MultihashSKEIN256200"
	IdentifierMultihashSKEIN256208	Identifier	=	"MultihashSKEIN256208"
	IdentifierMultihashSKEIN256216	Identifier	=	"MultihashSKEIN256216"
	IdentifierMultihashSKEIN256224	Identifier	=	"MultihashSKEIN256224"
	IdentifierMultihashSKEIN256232	Identifier	=	"MultihashSKEIN256232"
	IdentifierMultihashSKEIN256240	Identifier	=	"MultihashSKEIN256240"
	IdentifierMultihashSKEIN256248	Identifier	=	"MultihashSKEIN256248"
	IdentifierMultihashSKEIN256256	Identifier	=	"MultihashSKEIN256256"
	IdentifierMultihashSKEIN5128	Identifier	=	"MultihashSKEIN5128"
	IdentifierMultihashSKEIN51216	Identifier	=	"MultihashSKEIN51216"
	IdentifierMultihashSKEIN51224	Identifier	=	"MultihashSKEIN51224"
	IdentifierMultihashSKEIN51232	Identifier	=	"MultihashSKEIN51232"
	IdentifierMultihashSKEIN51240	Identifier	=	"MultihashSKEIN51240"
	IdentifierMultihashSKEIN51248	Identifier	=	"MultihashSKEIN51248"
	IdentifierMultihashSKEIN51256	Identifier	=	"MultihashSKEIN51256"
	IdentifierMultihashSKEIN51264	Identifier	=	"MultihashSKEIN51264"
	IdentifierMultihashSKEIN51272	Identifier	=	"MultihashSKEIN51272"
	IdentifierMultihashSKEIN51280	Identifier	=	"MultihashSKEIN51280"
	IdentifierMultihashSKEIN51288	Identifier	=	"MultihashSKEIN51288"
	IdentifierMultihashSKEIN51296	Identifier	=	"MultihashSKEIN51296"
	IdentifierMultihashSKEIN512104	Identifier	=	"MultihashSKEIN512104"
	IdentifierMultihashSKEIN512112	Identifier	=	"MultihashSKEIN512112"
	IdentifierMultihashSKEIN512120	Identifier	=	"MultihashSKEIN512120"
	IdentifierMultihashSKEIN512128	Identifier	=	"MultihashSKEIN512128"
	IdentifierMultihashSKEIN512136	Identifier	=	"MultihashSKEIN512136"
	IdentifierMultihashSKEIN512144	Identifier	=	"MultihashSKEIN512144"
	IdentifierMultihashSKEIN512152	Identifier	=	"MultihashSKEIN512152"
	IdentifierMultihashSKEIN512160	Identifier	=	"MultihashSKEIN512160"
	IdentifierMultihashSKEIN512168	Identifier	=	"MultihashSKEIN512168"
	IdentifierMultihashSKEIN512176	Identifier	=	"MultihashSKEIN512176"
	IdentifierMultihashSKEIN512184	Identifier	=	"MultihashSKEIN512184"
	IdentifierMultihashSKEIN512192	Identifier	=	"MultihashSKEIN512192"
	IdentifierMultihashSKEIN512200	Identifier	=	"MultihashSKEIN512200"
	IdentifierMultihashSKEIN512208	Identifier	=	"MultihashSKEIN512208"
	IdentifierMultihashSKEIN512216	Identifier	=	"MultihashSKEIN512216"
	IdentifierMultihashSKEIN512224	Identifier	=	"MultihashSKEIN512224"
	IdentifierMultihashSKEIN512232	Identifier	=	"MultihashSKEIN512232"
	IdentifierMultihashSKEIN512240	Identifier	=	"MultihashSKEIN512240"
	IdentifierMultihashSKEIN512248	Identifier	=	"MultihashSKEIN512248"
	IdentifierMultihashSKEIN512256	Identifier	=	"MultihashSKEIN512256"
	IdentifierMultihashSKEIN512264	Identifier	=	"MultihashSKEIN512264"
	IdentifierMultihashSKEIN512272	Identifier	=	"MultihashSKEIN512272"
	IdentifierMultihashSKEIN512280	Identifier	=	"MultihashSKEIN512280"
	IdentifierMultihashSKEIN512288	Identifier	=	"MultihashSKEIN512288"
	IdentifierMultihashSKEIN512296	Identifier	=	"MultihashSKEIN512296"
	IdentifierMultihashSKEIN512304	Identifier	=	"MultihashSKEIN512304"
	IdentifierMultihashSKEIN512312	Identifier	=	"MultihashSKEIN512312"
	IdentifierMultihashSKEIN512320	Identifier	=	"MultihashSKEIN512320"
	IdentifierMultihashSKEIN512328	Identifier	=	"MultihashSKEIN512328"
	IdentifierMultihashSKEIN512336	Identifier	=	"MultihashSKEIN512336"
	IdentifierMultihashSKEIN512344	Identifier	=	"MultihashSKEIN512344"
	IdentifierMultihashSKEIN512352	Identifier	=	"MultihashSKEIN512352"
	IdentifierMultihashSKEIN512360	Identifier	=	"MultihashSKEIN512360"
	IdentifierMultihashSKEIN512368	Identifier	=	"MultihashSKEIN512368"
	IdentifierMultihashSKEIN512376	Identifier	=	"MultihashSKEIN512376"
	IdentifierMultihashSKEIN512384	Identifier	=	"MultihashSKEIN512384"
	IdentifierMultihashSKEIN512392	Identifier	=	"MultihashSKEIN512392"
	IdentifierMultihashSKEIN512400	Identifier	=	"MultihashSKEIN512400"
	IdentifierMultihashSKEIN512408	Identifier	=	"MultihashSKEIN512408"
	IdentifierMultihashSKEIN512416	Identifier	=	"MultihashSKEIN512416"
	IdentifierMultihashSKEIN512424	Identifier	=	"MultihashSKEIN512424"
	IdentifierMultihashSKEIN512432	Identifier	=	"MultihashSKEIN512432"
	IdentifierMultihashSKEIN512440	Identifier	=	"MultihashSKEIN512440"
	IdentifierMultihashSKEIN512448	Identifier	=	"MultihashSKEIN512448"
	IdentifierMultihashSKEIN512456	Identifier	=	"MultihashSKEIN512456"
	IdentifierMultihashSKEIN512464	Identifier	=	"MultihashSKEIN512464"
	IdentifierMultihashSKEIN512472	Identifier	=	"MultihashSKEIN512472"
	IdentifierMultihashSKEIN512480	Identifier	=	"MultihashSKEIN512480"
	IdentifierMultihashSKEIN512488	Identifier	=	"MultihashSKEIN512488"
	IdentifierMultihashSKEIN512496	Identifier	=	"MultihashSKEIN512496"
	IdentifierMultihashSKEIN512504	Identifier	=	"MultihashSKEIN512504"
	IdentifierMultihashSKEIN512512	Identifier	=	"MultihashSKEIN512512"
	IdentifierMultihashSKEIN10248	Identifier	=	"MultihashSKEIN10248"
	IdentifierMultihashSKEIN102416	Identifier	=	"MultihashSKEIN102416"
	IdentifierMultihashSKEIN102424	Identifier	=	"MultihashSKEIN102424"
	IdentifierMultihashSKEIN102432	Identifier	=	"MultihashSKEIN102432"
	IdentifierMultihashSKEIN102440	Identifier	=	"MultihashSKEIN102440"
	IdentifierMultihashSKEIN102448	Identifier	=	"MultihashSKEIN102448"
	IdentifierMultihashSKEIN102456	Identifier	=	"MultihashSKEIN102456"
	IdentifierMultihashSKEIN102464	Identifier	=	"MultihashSKEIN102464"
	IdentifierMultihashSKEIN102472	Identifier	=	"MultihashSKEIN102472"
	IdentifierMultihashSKEIN102480	Identifier	=	"MultihashSKEIN102480"
	IdentifierMultihashSKEIN102488	Identifier	=	"MultihashSKEIN102488"
	IdentifierMultihashSKEIN102496	Identifier	=	"MultihashSKEIN102496"
	IdentifierMultihashSKEIN1024104	Identifier	=	"MultihashSKEIN1024104"
	IdentifierMultihashSKEIN1024112	Identifier	=	"MultihashSKEIN1024112"
	IdentifierMultihashSKEIN1024120	Identifier	=	"MultihashSKEIN1024120"
	IdentifierMultihashSKEIN1024128	Identifier	=	"MultihashSKEIN1024128"
	IdentifierMultihashSKEIN1024136	Identifier	=	"MultihashSKEIN1024136"
	IdentifierMultihashSKEIN1024144	Identifier	=	"MultihashSKEIN1024144"
	IdentifierMultihashSKEIN1024152	Identifier	=	"MultihashSKEIN1024152"
	IdentifierMultihashSKEIN1024160	Identifier	=	"MultihashSKEIN1024160"
	IdentifierMultihashSKEIN1024168	Identifier	=	"MultihashSKEIN1024168"
	IdentifierMultihashSKEIN1024176	Identifier	=	"MultihashSKEIN1024176"
	IdentifierMultihashSKEIN1024184	Identifier	=	"MultihashSKEIN1024184"
	IdentifierMultihashSKEIN1024192	Identifier	=	"MultihashSKEIN1024192"
	IdentifierMultihashSKEIN1024200	Identifier	=	"MultihashSKEIN1024200"
	IdentifierMultihashSKEIN1024208	Identifier	=	"MultihashSKEIN1024208"
	IdentifierMultihashSKEIN1024216	Identifier	=	"MultihashSKEIN1024216"
	IdentifierMultihashSKEIN1024224	Identifier	=	"MultihashSKEIN1024224"
	IdentifierMultihashSKEIN1024232	Identifier	=	"MultihashSKEIN1024232"
	IdentifierMultihashSKEIN1024240	Identifier	=	"MultihashSKEIN1024240"
	IdentifierMultihashSKEIN1024248	Identifier	=	"MultihashSKEIN1024248"
	IdentifierMultihashSKEIN1024256	Identifier	=	"MultihashSKEIN1024256"
	IdentifierMultihashSKEIN1024264	Identifier	=	"MultihashSKEIN1024264"
	IdentifierMultihashSKEIN1024272	Identifier	=	"MultihashSKEIN1024272"
	IdentifierMultihashSKEIN1024280	Identifier	=	"MultihashSKEIN1024280"
	IdentifierMultihashSKEIN1024288	Identifier	=	"MultihashSKEIN1024288"
	IdentifierMultihashSKEIN1024296	Identifier	=	"MultihashSKEIN1024296"
	IdentifierMultihashSKEIN1024304	Identifier	=	"MultihashSKEIN1024304"
	IdentifierMultihashSKEIN1024312	Identifier	=	"MultihashSKEIN1024312"
	IdentifierMultihashSKEIN1024320	Identifier	=	"MultihashSKEIN1024320"
	IdentifierMultihashSKEIN1024328	Identifier	=	"MultihashSKEIN1024328"
	IdentifierMultihashSKEIN1024336	Identifier	=	"MultihashSKEIN1024336"
	IdentifierMultihashSKEIN1024344	Identifier	=	"MultihashSKEIN1024344"
	IdentifierMultihashSKEIN1024352	Identifier	=	"MultihashSKEIN1024352"
	IdentifierMultihashSKEIN1024360	Identifier	=	"MultihashSKEIN1024360"
	IdentifierMultihashSKEIN1024368	Identifier	=	"MultihashSKEIN1024368"
	IdentifierMultihashSKEIN1024376	Identifier	=	"MultihashSKEIN1024376"
	IdentifierMultihashSKEIN1024384	Identifier	=	"MultihashSKEIN1024384"
	IdentifierMultihashSKEIN1024392	Identifier	=	"MultihashSKEIN1024392"
	IdentifierMultihashSKEIN1024400	Identifier	=	"MultihashSKEIN1024400"
	IdentifierMultihashSKEIN1024408	Identifier	=	"MultihashSKEIN1024408"
	IdentifierMultihashSKEIN1024416	Identifier	=	"MultihashSKEIN1024416"
	IdentifierMultihashSKEIN1024424	Identifier	=	"MultihashSKEIN1024424"
	IdentifierMultihashSKEIN1024432	Identifier	=	"MultihashSKEIN1024432"
	IdentifierMultihashSKEIN1024440	Identifier	=	"MultihashSKEIN1024440"
	IdentifierMultihashSKEIN1024448	Identifier	=	"MultihashSKEIN1024448"
	IdentifierMultihashSKEIN1024456	Identifier	=	"MultihashSKEIN1024456"
	IdentifierMultihashSKEIN1024464	Identifier	=	"MultihashSKEIN1024464"
	IdentifierMultihashSKEIN1024472	Identifier	=	"MultihashSKEIN1024472"
	IdentifierMultihashSKEIN1024480	Identifier	=	"MultihashSKEIN1024480"
	IdentifierMultihashSKEIN1024488	Identifier	=	"MultihashSKEIN1024488"
	IdentifierMultihashSKEIN1024496	Identifier	=	"MultihashSKEIN1024496"
	IdentifierMultihashSKEIN1024504	Identifier	=	"MultihashSKEIN1024504"
	IdentifierMultihashSKEIN1024512	Identifier	=	"MultihashSKEIN1024512"
	IdentifierMultihashSKEIN1024520	Identifier	=	"MultihashSKEIN1024520"
	IdentifierMultihashSKEIN1024528	Identifier	=	"MultihashSKEIN1024528"
	IdentifierMultihashSKEIN1024536	Identifier	=	"MultihashSKEIN1024536"
	IdentifierMultihashSKEIN1024544	Identifier	=	"MultihashSKEIN1024544"
	IdentifierMultihashSKEIN1024552	Identifier	=	"MultihashSKEIN1024552"
	IdentifierMultihashSKEIN1024560	Identifier	=	"MultihashSKEIN1024560"
	IdentifierMultihashSKEIN1024568	Identifier	=	"MultihashSKEIN1024568"
	IdentifierMultihashSKEIN1024576	Identifier	=	"MultihashSKEIN1024576"
	IdentifierMultihashSKEIN1024584	Identifier	=	"MultihashSKEIN1024584"
	IdentifierMultihashSKEIN1024592	Identifier	=	"MultihashSKEIN1024592"
	IdentifierMultihashSKEIN1024600	Identifier	=	"MultihashSKEIN1024600"
	IdentifierMultihashSKEIN1024608	Identifier	=	"MultihashSKEIN1024608"
	IdentifierMultihashSKEIN1024616	Identifier	=	"MultihashSKEIN1024616"
	IdentifierMultihashSKEIN1024624	Identifier	=	"MultihashSKEIN1024624"
	IdentifierMultihashSKEIN1024632	Identifier	=	"MultihashSKEIN1024632"
	IdentifierMultihashSKEIN1024640	Identifier	=	"MultihashSKEIN1024640"
	IdentifierMultihashSKEIN1024648	Identifier	=	"MultihashSKEIN1024648"
	IdentifierMultihashSKEIN1024656	Identifier	=	"MultihashSKEIN1024656"
	IdentifierMultihashSKEIN1024664	Identifier	=	"MultihashSKEIN1024664"
	IdentifierMultihashSKEIN1024672	Identifier	=	"MultihashSKEIN1024672"
	IdentifierMultihashSKEIN1024680	Identifier	=	"MultihashSKEIN1024680"
	IdentifierMultihashSKEIN1024688	Identifier	=	"MultihashSKEIN1024688"
	IdentifierMultihashSKEIN1024696	Identifier	=	"MultihashSKEIN1024696"
	IdentifierMultihashSKEIN1024704	Identifier	=	"MultihashSKEIN1024704"
	IdentifierMultihashSKEIN1024712	Identifier	=	"MultihashSKEIN1024712"
	IdentifierMultihashSKEIN1024720	Identifier	=	"MultihashSKEIN1024720"
	IdentifierMultihashSKEIN1024728	Identifier	=	"MultihashSKEIN1024728"
	IdentifierMultihashSKEIN1024736	Identifier	=	"MultihashSKEIN1024736"
	IdentifierMultihashSKEIN1024744	Identifier	=	"MultihashSKEIN1024744"
	IdentifierMultihashSKEIN1024752	Identifier	=	"MultihashSKEIN1024752"
	IdentifierMultihashSKEIN1024760	Identifier	=	"MultihashSKEIN1024760"
	IdentifierMultihashSKEIN1024768	Identifier	=	"MultihashSKEIN1024768"
	IdentifierMultihashSKEIN1024776	Identifier	=	"MultihashSKEIN1024776"
	IdentifierMultihashSKEIN1024784	Identifier	=	"MultihashSKEIN1024784"
	IdentifierMultihashSKEIN1024792	Identifier	=	"MultihashSKEIN1024792"
	IdentifierMultihashSKEIN1024800	Identifier	=	"MultihashSKEIN1024800"
	IdentifierMultihashSKEIN1024808	Identifier	=	"MultihashSKEIN1024808"
	IdentifierMultihashSKEIN1024816	Identifier	=	"MultihashSKEIN1024816"
	IdentifierMultihashSKEIN1024824	Identifier	=	"MultihashSKEIN1024824"
	IdentifierMultihashSKEIN1024832	Identifier	=	"MultihashSKEIN1024832"
	IdentifierMultihashSKEIN1024840	Identifier	=	"MultihashSKEIN1024840"
	IdentifierMultihashSKEIN1024848	Identifier	=	"MultihashSKEIN1024848"
	IdentifierMultihashSKEIN1024856	Identifier	=	"MultihashSKEIN1024856"
	IdentifierMultihashSKEIN1024864	Identifier	=	"MultihashSKEIN1024864"
	IdentifierMultihashSKEIN1024872	Identifier	=	"MultihashSKEIN1024872"
	IdentifierMultihashSKEIN1024880	Identifier	=	"MultihashSKEIN1024880"
	IdentifierMultihashSKEIN1024888	Identifier	=	"MultihashSKEIN1024888"
	IdentifierMultihashSKEIN1024896	Identifier	=	"MultihashSKEIN1024896"
	IdentifierMultihashSKEIN1024904	Identifier	=	"MultihashSKEIN1024904"
	IdentifierMultihashSKEIN1024912	Identifier	=	"MultihashSKEIN1024912"
	IdentifierMultihashSKEIN1024920	Identifier	=	"MultihashSKEIN1024920"
	IdentifierMultihashSKEIN1024928	Identifier	=	"MultihashSKEIN1024928"
	IdentifierMultihashSKEIN1024936	Identifier	=	"MultihashSKEIN1024936"
	IdentifierMultihashSKEIN1024944	Identifier	=	"MultihashSKEIN1024944"
	IdentifierMultihashSKEIN1024952	Identifier	=	"MultihashSKEIN1024952"
	IdentifierMultihashSKEIN1024960	Identifier	=	"MultihashSKEIN1024960"
	IdentifierMultihashSKEIN1024968	Identifier	=	"MultihashSKEIN1024968"
	IdentifierMultihashSKEIN1024976	Identifier	=	"MultihashSKEIN1024976"
	IdentifierMultihashSKEIN1024984	Identifier	=	"MultihashSKEIN1024984"
	IdentifierMultihashSKEIN1024992	Identifier	=	"MultihashSKEIN1024992"
	IdentifierMultihashSKEIN10241000	Identifier	=	"MultihashSKEIN10241000"
	IdentifierMultihashSKEIN10241008	Identifier	=	"MultihashSKEIN10241008"
	IdentifierMultihashSKEIN10241016	Identifier	=	"MultihashSKEIN10241016"
	IdentifierMultihashSKEIN10241024	Identifier	=	"MultihashSKEIN10241024"
	IdentifierHashXXH32	Identifier	=	"HashXXH32"
	IdentifierHashXXH64	Identifier	=	"HashXXH64"
	IdentifierHashXXH364	Identifier	=	"HashXXH364"
	IdentifierHashXXH3128	Identifier	=	"HashXXH3128"
	IdentifierMultihashPOSEIDONBLS12381A2FC1	Identifier	=	"MultihashPOSEIDONBLS12381A2FC1"
	IdentifierMultihashPOSEIDONBLS12381A2FC1SC	Identifier	=	"MultihashPOSEIDONBLS12381A2FC1SC"
	IdentifierIpldRDFC1	Identifier	=	"IpldRDFC1"
	IdentifierSerializationSSZ	Identifier	=	"SerializationSSZ"
	IdentifierMultihashSSZSHA2256BMT	Identifier	=	"MultihashSSZSHA2256BMT"
	IdentifierIpldJSONJCS	Identifier	=	"IpldJSONJCS"
	IdentifierSofthashISCC	Identifier	=	"SofthashISCC"
	IdentifierZeroxcertZEROXCERTIMPRINT256	Identifier	=	"ZeroxcertZEROXCERTIMPRINT256"
	IdentifierVarsigNONSTANDARDSIG	Identifier	=	"VarsigNONSTANDARDSIG"
	IdentifierVarsigES256K	Identifier	=	"VarsigES256K"
	IdentifierVarsigBLS12381G1SIG	Identifier	=	"VarsigBLS12381G1SIG"
	IdentifierVarsigBLS12381G2SIG	Identifier	=	"VarsigBLS12381G2SIG"
	IdentifierVarsigEDDSA	Identifier	=	"VarsigEDDSA"
	IdentifierVarsigEIP191	Identifier	=	"VarsigEIP191"
	IdentifierKeyJWKJCSPUB	Identifier	=	"KeyJWKJCSPUB"
	IdentifierFilecoinFILCOMMITMENTUNSEALED	Identifier	=	"FilecoinFILCOMMITMENTUNSEALED"
	IdentifierFilecoinFILCOMMITMENTSEALED	Identifier	=	"FilecoinFILCOMMITMENTSEALED"
	IdentifierMultiaddrPLAINTEXTV2	Identifier	=	"MultiaddrPLAINTEXTV2"
	IdentifierHolochainHOLOCHAINADRV0	Identifier	=	"HolochainHOLOCHAINADRV0"
	IdentifierHolochainHOLOCHAINADRV1	Identifier	=	"HolochainHOLOCHAINADRV1"
	IdentifierHolochainHOLOCHAINKEYV0	Identifier	=	"HolochainHOLOCHAINKEYV0"
	IdentifierHolochainHOLOCHAINKEYV1	Identifier	=	"HolochainHOLOCHAINKEYV1"
	IdentifierHolochainHOLOCHAINSIGV0	Identifier	=	"HolochainHOLOCHAINSIGV0"
	IdentifierHolochainHOLOCHAINSIGV1	Identifier	=	"HolochainHOLOCHAINSIGV1"
	IdentifierNamespaceSKYNETNS	Identifier	=	"NamespaceSKYNETNS"
	IdentifierNamespaceARWEAVENS	Identifier	=	"NamespaceARWEAVENS"
	IdentifierNamespaceSUBSPACENS	Identifier	=	"NamespaceSUBSPACENS"
	IdentifierNamespaceKUMANDRANS	Identifier	=	"NamespaceKUMANDRANS"
	IdentifierVarsigES256	Identifier	=	"VarsigES256"
	IdentifierVarsigES284	Identifier	=	"VarsigES284"
	IdentifierVarsigES512	Identifier	=	"VarsigES512"
	IdentifierVarsigRS256	Identifier	=	"VarsigRS256"
	IdentifierMultiaddrSCION	Identifier	=	"MultiaddrSCION"


)
/*
 * Identify code.
 */
func (code Header) String() (string) {
	switch code {
	case HeaderMultihashIDENTITY:
		return string(IdentifierMultihashIDENTITY)
	case HeaderCidCIDV1:
		return string(IdentifierCidCIDV1)
	case HeaderCidCIDV2:
		return string(IdentifierCidCIDV2)
	case HeaderCidCIDV3:
		return string(IdentifierCidCIDV3)
	case HeaderMultiaddrIP4:
		return string(IdentifierMultiaddrIP4)
	case HeaderMultiaddrTCP:
		return string(IdentifierMultiaddrTCP)
	case HeaderMultihashSHA1:
		return string(IdentifierMultihashSHA1)
	case HeaderMultihashSHA2256:
		return string(IdentifierMultihashSHA2256)
	case HeaderMultihashSHA2512:
		return string(IdentifierMultihashSHA2512)
	case HeaderMultihashSHA3512:
		return string(IdentifierMultihashSHA3512)
	case HeaderMultihashSHA3384:
		return string(IdentifierMultihashSHA3384)
	case HeaderMultihashSHA3256:
		return string(IdentifierMultihashSHA3256)
	case HeaderMultihashSHA3224:
		return string(IdentifierMultihashSHA3224)
	case HeaderMultihashSHAKE128:
		return string(IdentifierMultihashSHAKE128)
	case HeaderMultihashSHAKE256:
		return string(IdentifierMultihashSHAKE256)
	case HeaderMultihashKECCAK224:
		return string(IdentifierMultihashKECCAK224)
	case HeaderMultihashKECCAK256:
		return string(IdentifierMultihashKECCAK256)
	case HeaderMultihashKECCAK384:
		return string(IdentifierMultihashKECCAK384)
	case HeaderMultihashKECCAK512:
		return string(IdentifierMultihashKECCAK512)
	case HeaderMultihashBLAKE3:
		return string(IdentifierMultihashBLAKE3)
	case HeaderMultihashSHA2384:
		return string(IdentifierMultihashSHA2384)
	case HeaderMultiaddrDCCP:
		return string(IdentifierMultiaddrDCCP)
	case HeaderHashMURMUR3X6464:
		return string(IdentifierHashMURMUR3X6464)
	case HeaderHashMURMUR332:
		return string(IdentifierHashMURMUR332)
	case HeaderMultiaddrIP6:
		return string(IdentifierMultiaddrIP6)
	case HeaderMultiaddrIP6ZONE:
		return string(IdentifierMultiaddrIP6ZONE)
	case HeaderMultiaddrIPCIDR:
		return string(IdentifierMultiaddrIPCIDR)
	case HeaderNamespacePATH:
		return string(IdentifierNamespacePATH)
	case HeaderMultiformatMULTICODEC:
		return string(IdentifierMultiformatMULTICODEC)
	case HeaderMultiformatMULTIHASH:
		return string(IdentifierMultiformatMULTIHASH)
	case HeaderMultiformatMULTIADDR:
		return string(IdentifierMultiformatMULTIADDR)
	case HeaderMultiformatMULTIBASE:
		return string(IdentifierMultiformatMULTIBASE)
	case HeaderMultiformatVARSIG:
		return string(IdentifierMultiformatVARSIG)
	case HeaderMultiaddrDNS:
		return string(IdentifierMultiaddrDNS)
	case HeaderMultiaddrDNS4:
		return string(IdentifierMultiaddrDNS4)
	case HeaderMultiaddrDNS6:
		return string(IdentifierMultiaddrDNS6)
	case HeaderMultiaddrDNSADDR:
		return string(IdentifierMultiaddrDNSADDR)
	case HeaderSerializationPROTOBUF:
		return string(IdentifierSerializationPROTOBUF)
	case HeaderIpldCBOR:
		return string(IdentifierIpldCBOR)
	case HeaderIpldRAW:
		return string(IdentifierIpldRAW)
	case HeaderMultihashDBLSHA2256:
		return string(IdentifierMultihashDBLSHA2256)
	case HeaderSerializationRLP:
		return string(IdentifierSerializationRLP)
	case HeaderSerializationBENCODE:
		return string(IdentifierSerializationBENCODE)
	case HeaderIpldDAGPB:
		return string(IdentifierIpldDAGPB)
	case HeaderIpldDAGCBOR:
		return string(IdentifierIpldDAGCBOR)
	case HeaderIpldLIBP2PKEY:
		return string(IdentifierIpldLIBP2PKEY)
	case HeaderIpldGITRAW:
		return string(IdentifierIpldGITRAW)
	case HeaderIpldTORRENTINFO:
		return string(IdentifierIpldTORRENTINFO)
	case HeaderIpldTORRENTFILE:
		return string(IdentifierIpldTORRENTFILE)
	case HeaderIpldLEOFCOINBLOCK:
		return string(IdentifierIpldLEOFCOINBLOCK)
	case HeaderIpldLEOFCOINTX:
		return string(IdentifierIpldLEOFCOINTX)
	case HeaderIpldLEOFCOINPR:
		return string(IdentifierIpldLEOFCOINPR)
	case HeaderMultiaddrSCTP:
		return string(IdentifierMultiaddrSCTP)
	case HeaderIpldDAGJOSE:
		return string(IdentifierIpldDAGJOSE)
	case HeaderIpldDAGCOSE:
		return string(IdentifierIpldDAGCOSE)
	case HeaderNamespaceLBRY:
		return string(IdentifierNamespaceLBRY)
	case HeaderIpldETHBLOCK:
		return string(IdentifierIpldETHBLOCK)
	case HeaderIpldETHBLOCKLIST:
		return string(IdentifierIpldETHBLOCKLIST)
	case HeaderIpldETHTXTRIE:
		return string(IdentifierIpldETHTXTRIE)
	case HeaderIpldETHTX:
		return string(IdentifierIpldETHTX)
	case HeaderIpldETHTXRECEIPTTRIE:
		return string(IdentifierIpldETHTXRECEIPTTRIE)
	case HeaderIpldETHTXRECEIPT:
		return string(IdentifierIpldETHTXRECEIPT)
	case HeaderIpldETHSTATETRIE:
		return string(IdentifierIpldETHSTATETRIE)
	case HeaderIpldETHACCOUNTSNAPSHOT:
		return string(IdentifierIpldETHACCOUNTSNAPSHOT)
	case HeaderIpldETHSTORAGETRIE:
		return string(IdentifierIpldETHSTORAGETRIE)
	case HeaderIpldETHRECEIPTLOGTRIE:
		return string(IdentifierIpldETHRECEIPTLOGTRIE)
	case HeaderIpldETHRECEIPTLOG:
		return string(IdentifierIpldETHRECEIPTLOG)
	case HeaderKeyAES128:
		return string(IdentifierKeyAES128)
	case HeaderKeyAES192:
		return string(IdentifierKeyAES192)
	case HeaderKeyAES256:
		return string(IdentifierKeyAES256)
	case HeaderKeyCHACHA128:
		return string(IdentifierKeyCHACHA128)
	case HeaderKeyCHACHA256:
		return string(IdentifierKeyCHACHA256)
	case HeaderIpldBITCOINBLOCK:
		return string(IdentifierIpldBITCOINBLOCK)
	case HeaderIpldBITCOINTX:
		return string(IdentifierIpldBITCOINTX)
	case HeaderIpldBITCOINWITNESSCOMMITMENT:
		return string(IdentifierIpldBITCOINWITNESSCOMMITMENT)
	case HeaderIpldZCASHBLOCK:
		return string(IdentifierIpldZCASHBLOCK)
	case HeaderIpldZCASHTX:
		return string(IdentifierIpldZCASHTX)
	case HeaderMultiformatCAIP50:
		return string(IdentifierMultiformatCAIP50)
	case HeaderNamespaceSTREAMID:
		return string(IdentifierNamespaceSTREAMID)
	case HeaderIpldSTELLARBLOCK:
		return string(IdentifierIpldSTELLARBLOCK)
	case HeaderIpldSTELLARTX:
		return string(IdentifierIpldSTELLARTX)
	case HeaderMultihashMD4:
		return string(IdentifierMultihashMD4)
	case HeaderMultihashMD5:
		return string(IdentifierMultihashMD5)
	case HeaderIpldDECREDBLOCK:
		return string(IdentifierIpldDECREDBLOCK)
	case HeaderIpldDECREDTX:
		return string(IdentifierIpldDECREDTX)
	case HeaderNamespaceIPLD:
		return string(IdentifierNamespaceIPLD)
	case HeaderNamespaceIPFS:
		return string(IdentifierNamespaceIPFS)
	case HeaderNamespaceSWARM:
		return string(IdentifierNamespaceSWARM)
	case HeaderNamespaceIPNS:
		return string(IdentifierNamespaceIPNS)
	case HeaderNamespaceZERONET:
		return string(IdentifierNamespaceZERONET)
	case HeaderKeySECP256K1PUB:
		return string(IdentifierKeySECP256K1PUB)
	case HeaderNamespaceDNSLINK:
		return string(IdentifierNamespaceDNSLINK)
	case HeaderKeyBLS12381G1PUB:
		return string(IdentifierKeyBLS12381G1PUB)
	case HeaderKeyBLS12381G2PUB:
		return string(IdentifierKeyBLS12381G2PUB)
	case HeaderKeyX25519PUB:
		return string(IdentifierKeyX25519PUB)
	case HeaderKeyED25519PUB:
		return string(IdentifierKeyED25519PUB)
	case HeaderKeyBLS12381G1G2PUB:
		return string(IdentifierKeyBLS12381G1G2PUB)
	case HeaderKeySR25519PUB:
		return string(IdentifierKeySR25519PUB)
	case HeaderIpldDASHBLOCK:
		return string(IdentifierIpldDASHBLOCK)
	case HeaderIpldDASHTX:
		return string(IdentifierIpldDASHTX)
	case HeaderIpldSWARMMANIFEST:
		return string(IdentifierIpldSWARMMANIFEST)
	case HeaderIpldSWARMFEED:
		return string(IdentifierIpldSWARMFEED)
	case HeaderIpldBEESON:
		return string(IdentifierIpldBEESON)
	case HeaderMultiaddrUDP:
		return string(IdentifierMultiaddrUDP)
	case HeaderMultiaddrP2PWEBRTCSTAR:
		return string(IdentifierMultiaddrP2PWEBRTCSTAR)
	case HeaderMultiaddrP2PWEBRTCDIRECT:
		return string(IdentifierMultiaddrP2PWEBRTCDIRECT)
	case HeaderMultiaddrP2PSTARDUST:
		return string(IdentifierMultiaddrP2PSTARDUST)
	case HeaderMultiaddrWEBRTCDIRECT:
		return string(IdentifierMultiaddrWEBRTCDIRECT)
	case HeaderMultiaddrWEBRTC:
		return string(IdentifierMultiaddrWEBRTC)
	case HeaderMultiaddrP2PCIRCUIT:
		return string(IdentifierMultiaddrP2PCIRCUIT)
	case HeaderIpldDAGJSON:
		return string(IdentifierIpldDAGJSON)
	case HeaderMultiaddrUDT:
		return string(IdentifierMultiaddrUDT)
	case HeaderMultiaddrUTP:
		return string(IdentifierMultiaddrUTP)
	case HeaderHashCRC32:
		return string(IdentifierHashCRC32)
	case HeaderHashCRC64ECMA:
		return string(IdentifierHashCRC64ECMA)
	case HeaderMultiaddrUNIX:
		return string(IdentifierMultiaddrUNIX)
	case HeaderMultiaddrTHREAD:
		return string(IdentifierMultiaddrTHREAD)
	case HeaderMultiaddrP2P:
		return string(IdentifierMultiaddrP2P)
	case HeaderMultiaddrHTTPS:
		return string(IdentifierMultiaddrHTTPS)
	case HeaderMultiaddrONION:
		return string(IdentifierMultiaddrONION)
	case HeaderMultiaddrONION3:
		return string(IdentifierMultiaddrONION3)
	case HeaderMultiaddrGARLIC64:
		return string(IdentifierMultiaddrGARLIC64)
	case HeaderMultiaddrGARLIC32:
		return string(IdentifierMultiaddrGARLIC32)
	case HeaderMultiaddrTLS:
		return string(IdentifierMultiaddrTLS)
	case HeaderMultiaddrSNI:
		return string(IdentifierMultiaddrSNI)
	case HeaderMultiaddrNOISE:
		return string(IdentifierMultiaddrNOISE)
	case HeaderMultiaddrSHS:
		return string(IdentifierMultiaddrSHS)
	case HeaderMultiaddrQUIC:
		return string(IdentifierMultiaddrQUIC)
	case HeaderMultiaddrQUICV1:
		return string(IdentifierMultiaddrQUICV1)
	case HeaderMultiaddrWEBTRANSPORT:
		return string(IdentifierMultiaddrWEBTRANSPORT)
	case HeaderMultiaddrCERTHASH:
		return string(IdentifierMultiaddrCERTHASH)
	case HeaderMultiaddrWS:
		return string(IdentifierMultiaddrWS)
	case HeaderMultiaddrWSS:
		return string(IdentifierMultiaddrWSS)
	case HeaderMultiaddrP2PWEBSOCKETSTAR:
		return string(IdentifierMultiaddrP2PWEBSOCKETSTAR)
	case HeaderMultiaddrHTTP:
		return string(IdentifierMultiaddrHTTP)
	case HeaderIpldSWHID1SNP:
		return string(IdentifierIpldSWHID1SNP)
	case HeaderIpldJSON:
		return string(IdentifierIpldJSON)
	case HeaderSerializationMESSAGEPACK:
		return string(IdentifierSerializationMESSAGEPACK)
	case HeaderSerializationCAR:
		return string(IdentifierSerializationCAR)
	case HeaderSerializationIPNSRECORD:
		return string(IdentifierSerializationIPNSRECORD)
	case HeaderLibp2pLIBP2PPEERRECORD:
		return string(IdentifierLibp2pLIBP2PPEERRECORD)
	case HeaderLibp2pLIBP2PRELAYRSVP:
		return string(IdentifierLibp2pLIBP2PRELAYRSVP)
	case HeaderLibp2pMEMORYTRANSPORT:
		return string(IdentifierLibp2pMEMORYTRANSPORT)
	case HeaderSerializationCARINDEXSORTED:
		return string(IdentifierSerializationCARINDEXSORTED)
	case HeaderSerializationCARMULTIHASHINDEXSORTED:
		return string(IdentifierSerializationCARMULTIHASHINDEXSORTED)
	case HeaderTransportTRANSPORTBITSWAP:
		return string(IdentifierTransportTRANSPORTBITSWAP)
	case HeaderTransportTRANSPORTGRAPHSYNCFILECOINV1:
		return string(IdentifierTransportTRANSPORTGRAPHSYNCFILECOINV1)
	case HeaderTransportTRANSPORTIPFSGATEWAYHTTP:
		return string(IdentifierTransportTRANSPORTIPFSGATEWAYHTTP)
	case HeaderMultiformatMULTIDID:
		return string(IdentifierMultiformatMULTIDID)
	case HeaderMultihashSHA2256TRUNC254PADDED:
		return string(IdentifierMultihashSHA2256TRUNC254PADDED)
	case HeaderMultihashSHA2224:
		return string(IdentifierMultihashSHA2224)
	case HeaderMultihashSHA2512224:
		return string(IdentifierMultihashSHA2512224)
	case HeaderMultihashSHA2512256:
		return string(IdentifierMultihashSHA2512256)
	case HeaderHashMURMUR3X64128:
		return string(IdentifierHashMURMUR3X64128)
	case HeaderMultihashRIPEMD128:
		return string(IdentifierMultihashRIPEMD128)
	case HeaderMultihashRIPEMD160:
		return string(IdentifierMultihashRIPEMD160)
	case HeaderMultihashRIPEMD256:
		return string(IdentifierMultihashRIPEMD256)
	case HeaderMultihashRIPEMD320:
		return string(IdentifierMultihashRIPEMD320)
	case HeaderMultihashX11:
		return string(IdentifierMultihashX11)
	case HeaderKeyP256PUB:
		return string(IdentifierKeyP256PUB)
	case HeaderKeyP384PUB:
		return string(IdentifierKeyP384PUB)
	case HeaderKeyP521PUB:
		return string(IdentifierKeyP521PUB)
	case HeaderKeyED448PUB:
		return string(IdentifierKeyED448PUB)
	case HeaderKeyX448PUB:
		return string(IdentifierKeyX448PUB)
	case HeaderKeyRSAPUB:
		return string(IdentifierKeyRSAPUB)
	case HeaderKeySM2PUB:
		return string(IdentifierKeySM2PUB)
	case HeaderKeyED25519PRIV:
		return string(IdentifierKeyED25519PRIV)
	case HeaderKeySECP256K1PRIV:
		return string(IdentifierKeySECP256K1PRIV)
	case HeaderKeyX25519PRIV:
		return string(IdentifierKeyX25519PRIV)
	case HeaderKeySR25519PRIV:
		return string(IdentifierKeySR25519PRIV)
	case HeaderKeyRSAPRIV:
		return string(IdentifierKeyRSAPRIV)
	case HeaderKeyP256PRIV:
		return string(IdentifierKeyP256PRIV)
	case HeaderKeyP384PRIV:
		return string(IdentifierKeyP384PRIV)
	case HeaderKeyP521PRIV:
		return string(IdentifierKeyP521PRIV)
	case HeaderMultihashKANGAROOTWELVE:
		return string(IdentifierMultihashKANGAROOTWELVE)
	case HeaderEncryptionAESGCM256:
		return string(IdentifierEncryptionAESGCM256)
	case HeaderMultiaddrSILVERPINE:
		return string(IdentifierMultiaddrSILVERPINE)
	case HeaderMultihashSM3256:
		return string(IdentifierMultihashSM3256)
	case HeaderHashSHA256A:
		return string(IdentifierHashSHA256A)
	case HeaderMultihashBLAKE2B8:
		return string(IdentifierMultihashBLAKE2B8)
	case HeaderMultihashBLAKE2B16:
		return string(IdentifierMultihashBLAKE2B16)
	case HeaderMultihashBLAKE2B24:
		return string(IdentifierMultihashBLAKE2B24)
	case HeaderMultihashBLAKE2B32:
		return string(IdentifierMultihashBLAKE2B32)
	case HeaderMultihashBLAKE2B40:
		return string(IdentifierMultihashBLAKE2B40)
	case HeaderMultihashBLAKE2B48:
		return string(IdentifierMultihashBLAKE2B48)
	case HeaderMultihashBLAKE2B56:
		return string(IdentifierMultihashBLAKE2B56)
	case HeaderMultihashBLAKE2B64:
		return string(IdentifierMultihashBLAKE2B64)
	case HeaderMultihashBLAKE2B72:
		return string(IdentifierMultihashBLAKE2B72)
	case HeaderMultihashBLAKE2B80:
		return string(IdentifierMultihashBLAKE2B80)
	case HeaderMultihashBLAKE2B88:
		return string(IdentifierMultihashBLAKE2B88)
	case HeaderMultihashBLAKE2B96:
		return string(IdentifierMultihashBLAKE2B96)
	case HeaderMultihashBLAKE2B104:
		return string(IdentifierMultihashBLAKE2B104)
	case HeaderMultihashBLAKE2B112:
		return string(IdentifierMultihashBLAKE2B112)
	case HeaderMultihashBLAKE2B120:
		return string(IdentifierMultihashBLAKE2B120)
	case HeaderMultihashBLAKE2B128:
		return string(IdentifierMultihashBLAKE2B128)
	case HeaderMultihashBLAKE2B136:
		return string(IdentifierMultihashBLAKE2B136)
	case HeaderMultihashBLAKE2B144:
		return string(IdentifierMultihashBLAKE2B144)
	case HeaderMultihashBLAKE2B152:
		return string(IdentifierMultihashBLAKE2B152)
	case HeaderMultihashBLAKE2B160:
		return string(IdentifierMultihashBLAKE2B160)
	case HeaderMultihashBLAKE2B168:
		return string(IdentifierMultihashBLAKE2B168)
	case HeaderMultihashBLAKE2B176:
		return string(IdentifierMultihashBLAKE2B176)
	case HeaderMultihashBLAKE2B184:
		return string(IdentifierMultihashBLAKE2B184)
	case HeaderMultihashBLAKE2B192:
		return string(IdentifierMultihashBLAKE2B192)
	case HeaderMultihashBLAKE2B200:
		return string(IdentifierMultihashBLAKE2B200)
	case HeaderMultihashBLAKE2B208:
		return string(IdentifierMultihashBLAKE2B208)
	case HeaderMultihashBLAKE2B216:
		return string(IdentifierMultihashBLAKE2B216)
	case HeaderMultihashBLAKE2B224:
		return string(IdentifierMultihashBLAKE2B224)
	case HeaderMultihashBLAKE2B232:
		return string(IdentifierMultihashBLAKE2B232)
	case HeaderMultihashBLAKE2B240:
		return string(IdentifierMultihashBLAKE2B240)
	case HeaderMultihashBLAKE2B248:
		return string(IdentifierMultihashBLAKE2B248)
	case HeaderMultihashBLAKE2B256:
		return string(IdentifierMultihashBLAKE2B256)
	case HeaderMultihashBLAKE2B264:
		return string(IdentifierMultihashBLAKE2B264)
	case HeaderMultihashBLAKE2B272:
		return string(IdentifierMultihashBLAKE2B272)
	case HeaderMultihashBLAKE2B280:
		return string(IdentifierMultihashBLAKE2B280)
	case HeaderMultihashBLAKE2B288:
		return string(IdentifierMultihashBLAKE2B288)
	case HeaderMultihashBLAKE2B296:
		return string(IdentifierMultihashBLAKE2B296)
	case HeaderMultihashBLAKE2B304:
		return string(IdentifierMultihashBLAKE2B304)
	case HeaderMultihashBLAKE2B312:
		return string(IdentifierMultihashBLAKE2B312)
	case HeaderMultihashBLAKE2B320:
		return string(IdentifierMultihashBLAKE2B320)
	case HeaderMultihashBLAKE2B328:
		return string(IdentifierMultihashBLAKE2B328)
	case HeaderMultihashBLAKE2B336:
		return string(IdentifierMultihashBLAKE2B336)
	case HeaderMultihashBLAKE2B344:
		return string(IdentifierMultihashBLAKE2B344)
	case HeaderMultihashBLAKE2B352:
		return string(IdentifierMultihashBLAKE2B352)
	case HeaderMultihashBLAKE2B360:
		return string(IdentifierMultihashBLAKE2B360)
	case HeaderMultihashBLAKE2B368:
		return string(IdentifierMultihashBLAKE2B368)
	case HeaderMultihashBLAKE2B376:
		return string(IdentifierMultihashBLAKE2B376)
	case HeaderMultihashBLAKE2B384:
		return string(IdentifierMultihashBLAKE2B384)
	case HeaderMultihashBLAKE2B392:
		return string(IdentifierMultihashBLAKE2B392)
	case HeaderMultihashBLAKE2B400:
		return string(IdentifierMultihashBLAKE2B400)
	case HeaderMultihashBLAKE2B408:
		return string(IdentifierMultihashBLAKE2B408)
	case HeaderMultihashBLAKE2B416:
		return string(IdentifierMultihashBLAKE2B416)
	case HeaderMultihashBLAKE2B424:
		return string(IdentifierMultihashBLAKE2B424)
	case HeaderMultihashBLAKE2B432:
		return string(IdentifierMultihashBLAKE2B432)
	case HeaderMultihashBLAKE2B440:
		return string(IdentifierMultihashBLAKE2B440)
	case HeaderMultihashBLAKE2B448:
		return string(IdentifierMultihashBLAKE2B448)
	case HeaderMultihashBLAKE2B456:
		return string(IdentifierMultihashBLAKE2B456)
	case HeaderMultihashBLAKE2B464:
		return string(IdentifierMultihashBLAKE2B464)
	case HeaderMultihashBLAKE2B472:
		return string(IdentifierMultihashBLAKE2B472)
	case HeaderMultihashBLAKE2B480:
		return string(IdentifierMultihashBLAKE2B480)
	case HeaderMultihashBLAKE2B488:
		return string(IdentifierMultihashBLAKE2B488)
	case HeaderMultihashBLAKE2B496:
		return string(IdentifierMultihashBLAKE2B496)
	case HeaderMultihashBLAKE2B504:
		return string(IdentifierMultihashBLAKE2B504)
	case HeaderMultihashBLAKE2B512:
		return string(IdentifierMultihashBLAKE2B512)
	case HeaderMultihashBLAKE2S8:
		return string(IdentifierMultihashBLAKE2S8)
	case HeaderMultihashBLAKE2S16:
		return string(IdentifierMultihashBLAKE2S16)
	case HeaderMultihashBLAKE2S24:
		return string(IdentifierMultihashBLAKE2S24)
	case HeaderMultihashBLAKE2S32:
		return string(IdentifierMultihashBLAKE2S32)
	case HeaderMultihashBLAKE2S40:
		return string(IdentifierMultihashBLAKE2S40)
	case HeaderMultihashBLAKE2S48:
		return string(IdentifierMultihashBLAKE2S48)
	case HeaderMultihashBLAKE2S56:
		return string(IdentifierMultihashBLAKE2S56)
	case HeaderMultihashBLAKE2S64:
		return string(IdentifierMultihashBLAKE2S64)
	case HeaderMultihashBLAKE2S72:
		return string(IdentifierMultihashBLAKE2S72)
	case HeaderMultihashBLAKE2S80:
		return string(IdentifierMultihashBLAKE2S80)
	case HeaderMultihashBLAKE2S88:
		return string(IdentifierMultihashBLAKE2S88)
	case HeaderMultihashBLAKE2S96:
		return string(IdentifierMultihashBLAKE2S96)
	case HeaderMultihashBLAKE2S104:
		return string(IdentifierMultihashBLAKE2S104)
	case HeaderMultihashBLAKE2S112:
		return string(IdentifierMultihashBLAKE2S112)
	case HeaderMultihashBLAKE2S120:
		return string(IdentifierMultihashBLAKE2S120)
	case HeaderMultihashBLAKE2S128:
		return string(IdentifierMultihashBLAKE2S128)
	case HeaderMultihashBLAKE2S136:
		return string(IdentifierMultihashBLAKE2S136)
	case HeaderMultihashBLAKE2S144:
		return string(IdentifierMultihashBLAKE2S144)
	case HeaderMultihashBLAKE2S152:
		return string(IdentifierMultihashBLAKE2S152)
	case HeaderMultihashBLAKE2S160:
		return string(IdentifierMultihashBLAKE2S160)
	case HeaderMultihashBLAKE2S168:
		return string(IdentifierMultihashBLAKE2S168)
	case HeaderMultihashBLAKE2S176:
		return string(IdentifierMultihashBLAKE2S176)
	case HeaderMultihashBLAKE2S184:
		return string(IdentifierMultihashBLAKE2S184)
	case HeaderMultihashBLAKE2S192:
		return string(IdentifierMultihashBLAKE2S192)
	case HeaderMultihashBLAKE2S200:
		return string(IdentifierMultihashBLAKE2S200)
	case HeaderMultihashBLAKE2S208:
		return string(IdentifierMultihashBLAKE2S208)
	case HeaderMultihashBLAKE2S216:
		return string(IdentifierMultihashBLAKE2S216)
	case HeaderMultihashBLAKE2S224:
		return string(IdentifierMultihashBLAKE2S224)
	case HeaderMultihashBLAKE2S232:
		return string(IdentifierMultihashBLAKE2S232)
	case HeaderMultihashBLAKE2S240:
		return string(IdentifierMultihashBLAKE2S240)
	case HeaderMultihashBLAKE2S248:
		return string(IdentifierMultihashBLAKE2S248)
	case HeaderMultihashBLAKE2S256:
		return string(IdentifierMultihashBLAKE2S256)
	case HeaderMultihashSKEIN2568:
		return string(IdentifierMultihashSKEIN2568)
	case HeaderMultihashSKEIN25616:
		return string(IdentifierMultihashSKEIN25616)
	case HeaderMultihashSKEIN25624:
		return string(IdentifierMultihashSKEIN25624)
	case HeaderMultihashSKEIN25632:
		return string(IdentifierMultihashSKEIN25632)
	case HeaderMultihashSKEIN25640:
		return string(IdentifierMultihashSKEIN25640)
	case HeaderMultihashSKEIN25648:
		return string(IdentifierMultihashSKEIN25648)
	case HeaderMultihashSKEIN25656:
		return string(IdentifierMultihashSKEIN25656)
	case HeaderMultihashSKEIN25664:
		return string(IdentifierMultihashSKEIN25664)
	case HeaderMultihashSKEIN25672:
		return string(IdentifierMultihashSKEIN25672)
	case HeaderMultihashSKEIN25680:
		return string(IdentifierMultihashSKEIN25680)
	case HeaderMultihashSKEIN25688:
		return string(IdentifierMultihashSKEIN25688)
	case HeaderMultihashSKEIN25696:
		return string(IdentifierMultihashSKEIN25696)
	case HeaderMultihashSKEIN256104:
		return string(IdentifierMultihashSKEIN256104)
	case HeaderMultihashSKEIN256112:
		return string(IdentifierMultihashSKEIN256112)
	case HeaderMultihashSKEIN256120:
		return string(IdentifierMultihashSKEIN256120)
	case HeaderMultihashSKEIN256128:
		return string(IdentifierMultihashSKEIN256128)
	case HeaderMultihashSKEIN256136:
		return string(IdentifierMultihashSKEIN256136)
	case HeaderMultihashSKEIN256144:
		return string(IdentifierMultihashSKEIN256144)
	case HeaderMultihashSKEIN256152:
		return string(IdentifierMultihashSKEIN256152)
	case HeaderMultihashSKEIN256160:
		return string(IdentifierMultihashSKEIN256160)
	case HeaderMultihashSKEIN256168:
		return string(IdentifierMultihashSKEIN256168)
	case HeaderMultihashSKEIN256176:
		return string(IdentifierMultihashSKEIN256176)
	case HeaderMultihashSKEIN256184:
		return string(IdentifierMultihashSKEIN256184)
	case HeaderMultihashSKEIN256192:
		return string(IdentifierMultihashSKEIN256192)
	case HeaderMultihashSKEIN256200:
		return string(IdentifierMultihashSKEIN256200)
	case HeaderMultihashSKEIN256208:
		return string(IdentifierMultihashSKEIN256208)
	case HeaderMultihashSKEIN256216:
		return string(IdentifierMultihashSKEIN256216)
	case HeaderMultihashSKEIN256224:
		return string(IdentifierMultihashSKEIN256224)
	case HeaderMultihashSKEIN256232:
		return string(IdentifierMultihashSKEIN256232)
	case HeaderMultihashSKEIN256240:
		return string(IdentifierMultihashSKEIN256240)
	case HeaderMultihashSKEIN256248:
		return string(IdentifierMultihashSKEIN256248)
	case HeaderMultihashSKEIN256256:
		return string(IdentifierMultihashSKEIN256256)
	case HeaderMultihashSKEIN5128:
		return string(IdentifierMultihashSKEIN5128)
	case HeaderMultihashSKEIN51216:
		return string(IdentifierMultihashSKEIN51216)
	case HeaderMultihashSKEIN51224:
		return string(IdentifierMultihashSKEIN51224)
	case HeaderMultihashSKEIN51232:
		return string(IdentifierMultihashSKEIN51232)
	case HeaderMultihashSKEIN51240:
		return string(IdentifierMultihashSKEIN51240)
	case HeaderMultihashSKEIN51248:
		return string(IdentifierMultihashSKEIN51248)
	case HeaderMultihashSKEIN51256:
		return string(IdentifierMultihashSKEIN51256)
	case HeaderMultihashSKEIN51264:
		return string(IdentifierMultihashSKEIN51264)
	case HeaderMultihashSKEIN51272:
		return string(IdentifierMultihashSKEIN51272)
	case HeaderMultihashSKEIN51280:
		return string(IdentifierMultihashSKEIN51280)
	case HeaderMultihashSKEIN51288:
		return string(IdentifierMultihashSKEIN51288)
	case HeaderMultihashSKEIN51296:
		return string(IdentifierMultihashSKEIN51296)
	case HeaderMultihashSKEIN512104:
		return string(IdentifierMultihashSKEIN512104)
	case HeaderMultihashSKEIN512112:
		return string(IdentifierMultihashSKEIN512112)
	case HeaderMultihashSKEIN512120:
		return string(IdentifierMultihashSKEIN512120)
	case HeaderMultihashSKEIN512128:
		return string(IdentifierMultihashSKEIN512128)
	case HeaderMultihashSKEIN512136:
		return string(IdentifierMultihashSKEIN512136)
	case HeaderMultihashSKEIN512144:
		return string(IdentifierMultihashSKEIN512144)
	case HeaderMultihashSKEIN512152:
		return string(IdentifierMultihashSKEIN512152)
	case HeaderMultihashSKEIN512160:
		return string(IdentifierMultihashSKEIN512160)
	case HeaderMultihashSKEIN512168:
		return string(IdentifierMultihashSKEIN512168)
	case HeaderMultihashSKEIN512176:
		return string(IdentifierMultihashSKEIN512176)
	case HeaderMultihashSKEIN512184:
		return string(IdentifierMultihashSKEIN512184)
	case HeaderMultihashSKEIN512192:
		return string(IdentifierMultihashSKEIN512192)
	case HeaderMultihashSKEIN512200:
		return string(IdentifierMultihashSKEIN512200)
	case HeaderMultihashSKEIN512208:
		return string(IdentifierMultihashSKEIN512208)
	case HeaderMultihashSKEIN512216:
		return string(IdentifierMultihashSKEIN512216)
	case HeaderMultihashSKEIN512224:
		return string(IdentifierMultihashSKEIN512224)
	case HeaderMultihashSKEIN512232:
		return string(IdentifierMultihashSKEIN512232)
	case HeaderMultihashSKEIN512240:
		return string(IdentifierMultihashSKEIN512240)
	case HeaderMultihashSKEIN512248:
		return string(IdentifierMultihashSKEIN512248)
	case HeaderMultihashSKEIN512256:
		return string(IdentifierMultihashSKEIN512256)
	case HeaderMultihashSKEIN512264:
		return string(IdentifierMultihashSKEIN512264)
	case HeaderMultihashSKEIN512272:
		return string(IdentifierMultihashSKEIN512272)
	case HeaderMultihashSKEIN512280:
		return string(IdentifierMultihashSKEIN512280)
	case HeaderMultihashSKEIN512288:
		return string(IdentifierMultihashSKEIN512288)
	case HeaderMultihashSKEIN512296:
		return string(IdentifierMultihashSKEIN512296)
	case HeaderMultihashSKEIN512304:
		return string(IdentifierMultihashSKEIN512304)
	case HeaderMultihashSKEIN512312:
		return string(IdentifierMultihashSKEIN512312)
	case HeaderMultihashSKEIN512320:
		return string(IdentifierMultihashSKEIN512320)
	case HeaderMultihashSKEIN512328:
		return string(IdentifierMultihashSKEIN512328)
	case HeaderMultihashSKEIN512336:
		return string(IdentifierMultihashSKEIN512336)
	case HeaderMultihashSKEIN512344:
		return string(IdentifierMultihashSKEIN512344)
	case HeaderMultihashSKEIN512352:
		return string(IdentifierMultihashSKEIN512352)
	case HeaderMultihashSKEIN512360:
		return string(IdentifierMultihashSKEIN512360)
	case HeaderMultihashSKEIN512368:
		return string(IdentifierMultihashSKEIN512368)
	case HeaderMultihashSKEIN512376:
		return string(IdentifierMultihashSKEIN512376)
	case HeaderMultihashSKEIN512384:
		return string(IdentifierMultihashSKEIN512384)
	case HeaderMultihashSKEIN512392:
		return string(IdentifierMultihashSKEIN512392)
	case HeaderMultihashSKEIN512400:
		return string(IdentifierMultihashSKEIN512400)
	case HeaderMultihashSKEIN512408:
		return string(IdentifierMultihashSKEIN512408)
	case HeaderMultihashSKEIN512416:
		return string(IdentifierMultihashSKEIN512416)
	case HeaderMultihashSKEIN512424:
		return string(IdentifierMultihashSKEIN512424)
	case HeaderMultihashSKEIN512432:
		return string(IdentifierMultihashSKEIN512432)
	case HeaderMultihashSKEIN512440:
		return string(IdentifierMultihashSKEIN512440)
	case HeaderMultihashSKEIN512448:
		return string(IdentifierMultihashSKEIN512448)
	case HeaderMultihashSKEIN512456:
		return string(IdentifierMultihashSKEIN512456)
	case HeaderMultihashSKEIN512464:
		return string(IdentifierMultihashSKEIN512464)
	case HeaderMultihashSKEIN512472:
		return string(IdentifierMultihashSKEIN512472)
	case HeaderMultihashSKEIN512480:
		return string(IdentifierMultihashSKEIN512480)
	case HeaderMultihashSKEIN512488:
		return string(IdentifierMultihashSKEIN512488)
	case HeaderMultihashSKEIN512496:
		return string(IdentifierMultihashSKEIN512496)
	case HeaderMultihashSKEIN512504:
		return string(IdentifierMultihashSKEIN512504)
	case HeaderMultihashSKEIN512512:
		return string(IdentifierMultihashSKEIN512512)
	case HeaderMultihashSKEIN10248:
		return string(IdentifierMultihashSKEIN10248)
	case HeaderMultihashSKEIN102416:
		return string(IdentifierMultihashSKEIN102416)
	case HeaderMultihashSKEIN102424:
		return string(IdentifierMultihashSKEIN102424)
	case HeaderMultihashSKEIN102432:
		return string(IdentifierMultihashSKEIN102432)
	case HeaderMultihashSKEIN102440:
		return string(IdentifierMultihashSKEIN102440)
	case HeaderMultihashSKEIN102448:
		return string(IdentifierMultihashSKEIN102448)
	case HeaderMultihashSKEIN102456:
		return string(IdentifierMultihashSKEIN102456)
	case HeaderMultihashSKEIN102464:
		return string(IdentifierMultihashSKEIN102464)
	case HeaderMultihashSKEIN102472:
		return string(IdentifierMultihashSKEIN102472)
	case HeaderMultihashSKEIN102480:
		return string(IdentifierMultihashSKEIN102480)
	case HeaderMultihashSKEIN102488:
		return string(IdentifierMultihashSKEIN102488)
	case HeaderMultihashSKEIN102496:
		return string(IdentifierMultihashSKEIN102496)
	case HeaderMultihashSKEIN1024104:
		return string(IdentifierMultihashSKEIN1024104)
	case HeaderMultihashSKEIN1024112:
		return string(IdentifierMultihashSKEIN1024112)
	case HeaderMultihashSKEIN1024120:
		return string(IdentifierMultihashSKEIN1024120)
	case HeaderMultihashSKEIN1024128:
		return string(IdentifierMultihashSKEIN1024128)
	case HeaderMultihashSKEIN1024136:
		return string(IdentifierMultihashSKEIN1024136)
	case HeaderMultihashSKEIN1024144:
		return string(IdentifierMultihashSKEIN1024144)
	case HeaderMultihashSKEIN1024152:
		return string(IdentifierMultihashSKEIN1024152)
	case HeaderMultihashSKEIN1024160:
		return string(IdentifierMultihashSKEIN1024160)
	case HeaderMultihashSKEIN1024168:
		return string(IdentifierMultihashSKEIN1024168)
	case HeaderMultihashSKEIN1024176:
		return string(IdentifierMultihashSKEIN1024176)
	case HeaderMultihashSKEIN1024184:
		return string(IdentifierMultihashSKEIN1024184)
	case HeaderMultihashSKEIN1024192:
		return string(IdentifierMultihashSKEIN1024192)
	case HeaderMultihashSKEIN1024200:
		return string(IdentifierMultihashSKEIN1024200)
	case HeaderMultihashSKEIN1024208:
		return string(IdentifierMultihashSKEIN1024208)
	case HeaderMultihashSKEIN1024216:
		return string(IdentifierMultihashSKEIN1024216)
	case HeaderMultihashSKEIN1024224:
		return string(IdentifierMultihashSKEIN1024224)
	case HeaderMultihashSKEIN1024232:
		return string(IdentifierMultihashSKEIN1024232)
	case HeaderMultihashSKEIN1024240:
		return string(IdentifierMultihashSKEIN1024240)
	case HeaderMultihashSKEIN1024248:
		return string(IdentifierMultihashSKEIN1024248)
	case HeaderMultihashSKEIN1024256:
		return string(IdentifierMultihashSKEIN1024256)
	case HeaderMultihashSKEIN1024264:
		return string(IdentifierMultihashSKEIN1024264)
	case HeaderMultihashSKEIN1024272:
		return string(IdentifierMultihashSKEIN1024272)
	case HeaderMultihashSKEIN1024280:
		return string(IdentifierMultihashSKEIN1024280)
	case HeaderMultihashSKEIN1024288:
		return string(IdentifierMultihashSKEIN1024288)
	case HeaderMultihashSKEIN1024296:
		return string(IdentifierMultihashSKEIN1024296)
	case HeaderMultihashSKEIN1024304:
		return string(IdentifierMultihashSKEIN1024304)
	case HeaderMultihashSKEIN1024312:
		return string(IdentifierMultihashSKEIN1024312)
	case HeaderMultihashSKEIN1024320:
		return string(IdentifierMultihashSKEIN1024320)
	case HeaderMultihashSKEIN1024328:
		return string(IdentifierMultihashSKEIN1024328)
	case HeaderMultihashSKEIN1024336:
		return string(IdentifierMultihashSKEIN1024336)
	case HeaderMultihashSKEIN1024344:
		return string(IdentifierMultihashSKEIN1024344)
	case HeaderMultihashSKEIN1024352:
		return string(IdentifierMultihashSKEIN1024352)
	case HeaderMultihashSKEIN1024360:
		return string(IdentifierMultihashSKEIN1024360)
	case HeaderMultihashSKEIN1024368:
		return string(IdentifierMultihashSKEIN1024368)
	case HeaderMultihashSKEIN1024376:
		return string(IdentifierMultihashSKEIN1024376)
	case HeaderMultihashSKEIN1024384:
		return string(IdentifierMultihashSKEIN1024384)
	case HeaderMultihashSKEIN1024392:
		return string(IdentifierMultihashSKEIN1024392)
	case HeaderMultihashSKEIN1024400:
		return string(IdentifierMultihashSKEIN1024400)
	case HeaderMultihashSKEIN1024408:
		return string(IdentifierMultihashSKEIN1024408)
	case HeaderMultihashSKEIN1024416:
		return string(IdentifierMultihashSKEIN1024416)
	case HeaderMultihashSKEIN1024424:
		return string(IdentifierMultihashSKEIN1024424)
	case HeaderMultihashSKEIN1024432:
		return string(IdentifierMultihashSKEIN1024432)
	case HeaderMultihashSKEIN1024440:
		return string(IdentifierMultihashSKEIN1024440)
	case HeaderMultihashSKEIN1024448:
		return string(IdentifierMultihashSKEIN1024448)
	case HeaderMultihashSKEIN1024456:
		return string(IdentifierMultihashSKEIN1024456)
	case HeaderMultihashSKEIN1024464:
		return string(IdentifierMultihashSKEIN1024464)
	case HeaderMultihashSKEIN1024472:
		return string(IdentifierMultihashSKEIN1024472)
	case HeaderMultihashSKEIN1024480:
		return string(IdentifierMultihashSKEIN1024480)
	case HeaderMultihashSKEIN1024488:
		return string(IdentifierMultihashSKEIN1024488)
	case HeaderMultihashSKEIN1024496:
		return string(IdentifierMultihashSKEIN1024496)
	case HeaderMultihashSKEIN1024504:
		return string(IdentifierMultihashSKEIN1024504)
	case HeaderMultihashSKEIN1024512:
		return string(IdentifierMultihashSKEIN1024512)
	case HeaderMultihashSKEIN1024520:
		return string(IdentifierMultihashSKEIN1024520)
	case HeaderMultihashSKEIN1024528:
		return string(IdentifierMultihashSKEIN1024528)
	case HeaderMultihashSKEIN1024536:
		return string(IdentifierMultihashSKEIN1024536)
	case HeaderMultihashSKEIN1024544:
		return string(IdentifierMultihashSKEIN1024544)
	case HeaderMultihashSKEIN1024552:
		return string(IdentifierMultihashSKEIN1024552)
	case HeaderMultihashSKEIN1024560:
		return string(IdentifierMultihashSKEIN1024560)
	case HeaderMultihashSKEIN1024568:
		return string(IdentifierMultihashSKEIN1024568)
	case HeaderMultihashSKEIN1024576:
		return string(IdentifierMultihashSKEIN1024576)
	case HeaderMultihashSKEIN1024584:
		return string(IdentifierMultihashSKEIN1024584)
	case HeaderMultihashSKEIN1024592:
		return string(IdentifierMultihashSKEIN1024592)
	case HeaderMultihashSKEIN1024600:
		return string(IdentifierMultihashSKEIN1024600)
	case HeaderMultihashSKEIN1024608:
		return string(IdentifierMultihashSKEIN1024608)
	case HeaderMultihashSKEIN1024616:
		return string(IdentifierMultihashSKEIN1024616)
	case HeaderMultihashSKEIN1024624:
		return string(IdentifierMultihashSKEIN1024624)
	case HeaderMultihashSKEIN1024632:
		return string(IdentifierMultihashSKEIN1024632)
	case HeaderMultihashSKEIN1024640:
		return string(IdentifierMultihashSKEIN1024640)
	case HeaderMultihashSKEIN1024648:
		return string(IdentifierMultihashSKEIN1024648)
	case HeaderMultihashSKEIN1024656:
		return string(IdentifierMultihashSKEIN1024656)
	case HeaderMultihashSKEIN1024664:
		return string(IdentifierMultihashSKEIN1024664)
	case HeaderMultihashSKEIN1024672:
		return string(IdentifierMultihashSKEIN1024672)
	case HeaderMultihashSKEIN1024680:
		return string(IdentifierMultihashSKEIN1024680)
	case HeaderMultihashSKEIN1024688:
		return string(IdentifierMultihashSKEIN1024688)
	case HeaderMultihashSKEIN1024696:
		return string(IdentifierMultihashSKEIN1024696)
	case HeaderMultihashSKEIN1024704:
		return string(IdentifierMultihashSKEIN1024704)
	case HeaderMultihashSKEIN1024712:
		return string(IdentifierMultihashSKEIN1024712)
	case HeaderMultihashSKEIN1024720:
		return string(IdentifierMultihashSKEIN1024720)
	case HeaderMultihashSKEIN1024728:
		return string(IdentifierMultihashSKEIN1024728)
	case HeaderMultihashSKEIN1024736:
		return string(IdentifierMultihashSKEIN1024736)
	case HeaderMultihashSKEIN1024744:
		return string(IdentifierMultihashSKEIN1024744)
	case HeaderMultihashSKEIN1024752:
		return string(IdentifierMultihashSKEIN1024752)
	case HeaderMultihashSKEIN1024760:
		return string(IdentifierMultihashSKEIN1024760)
	case HeaderMultihashSKEIN1024768:
		return string(IdentifierMultihashSKEIN1024768)
	case HeaderMultihashSKEIN1024776:
		return string(IdentifierMultihashSKEIN1024776)
	case HeaderMultihashSKEIN1024784:
		return string(IdentifierMultihashSKEIN1024784)
	case HeaderMultihashSKEIN1024792:
		return string(IdentifierMultihashSKEIN1024792)
	case HeaderMultihashSKEIN1024800:
		return string(IdentifierMultihashSKEIN1024800)
	case HeaderMultihashSKEIN1024808:
		return string(IdentifierMultihashSKEIN1024808)
	case HeaderMultihashSKEIN1024816:
		return string(IdentifierMultihashSKEIN1024816)
	case HeaderMultihashSKEIN1024824:
		return string(IdentifierMultihashSKEIN1024824)
	case HeaderMultihashSKEIN1024832:
		return string(IdentifierMultihashSKEIN1024832)
	case HeaderMultihashSKEIN1024840:
		return string(IdentifierMultihashSKEIN1024840)
	case HeaderMultihashSKEIN1024848:
		return string(IdentifierMultihashSKEIN1024848)
	case HeaderMultihashSKEIN1024856:
		return string(IdentifierMultihashSKEIN1024856)
	case HeaderMultihashSKEIN1024864:
		return string(IdentifierMultihashSKEIN1024864)
	case HeaderMultihashSKEIN1024872:
		return string(IdentifierMultihashSKEIN1024872)
	case HeaderMultihashSKEIN1024880:
		return string(IdentifierMultihashSKEIN1024880)
	case HeaderMultihashSKEIN1024888:
		return string(IdentifierMultihashSKEIN1024888)
	case HeaderMultihashSKEIN1024896:
		return string(IdentifierMultihashSKEIN1024896)
	case HeaderMultihashSKEIN1024904:
		return string(IdentifierMultihashSKEIN1024904)
	case HeaderMultihashSKEIN1024912:
		return string(IdentifierMultihashSKEIN1024912)
	case HeaderMultihashSKEIN1024920:
		return string(IdentifierMultihashSKEIN1024920)
	case HeaderMultihashSKEIN1024928:
		return string(IdentifierMultihashSKEIN1024928)
	case HeaderMultihashSKEIN1024936:
		return string(IdentifierMultihashSKEIN1024936)
	case HeaderMultihashSKEIN1024944:
		return string(IdentifierMultihashSKEIN1024944)
	case HeaderMultihashSKEIN1024952:
		return string(IdentifierMultihashSKEIN1024952)
	case HeaderMultihashSKEIN1024960:
		return string(IdentifierMultihashSKEIN1024960)
	case HeaderMultihashSKEIN1024968:
		return string(IdentifierMultihashSKEIN1024968)
	case HeaderMultihashSKEIN1024976:
		return string(IdentifierMultihashSKEIN1024976)
	case HeaderMultihashSKEIN1024984:
		return string(IdentifierMultihashSKEIN1024984)
	case HeaderMultihashSKEIN1024992:
		return string(IdentifierMultihashSKEIN1024992)
	case HeaderMultihashSKEIN10241000:
		return string(IdentifierMultihashSKEIN10241000)
	case HeaderMultihashSKEIN10241008:
		return string(IdentifierMultihashSKEIN10241008)
	case HeaderMultihashSKEIN10241016:
		return string(IdentifierMultihashSKEIN10241016)
	case HeaderMultihashSKEIN10241024:
		return string(IdentifierMultihashSKEIN10241024)
	case HeaderHashXXH32:
		return string(IdentifierHashXXH32)
	case HeaderHashXXH64:
		return string(IdentifierHashXXH64)
	case HeaderHashXXH364:
		return string(IdentifierHashXXH364)
	case HeaderHashXXH3128:
		return string(IdentifierHashXXH3128)
	case HeaderMultihashPOSEIDONBLS12381A2FC1:
		return string(IdentifierMultihashPOSEIDONBLS12381A2FC1)
	case HeaderMultihashPOSEIDONBLS12381A2FC1SC:
		return string(IdentifierMultihashPOSEIDONBLS12381A2FC1SC)
	case HeaderIpldRDFC1:
		return string(IdentifierIpldRDFC1)
	case HeaderSerializationSSZ:
		return string(IdentifierSerializationSSZ)
	case HeaderMultihashSSZSHA2256BMT:
		return string(IdentifierMultihashSSZSHA2256BMT)
	case HeaderIpldJSONJCS:
		return string(IdentifierIpldJSONJCS)
	case HeaderSofthashISCC:
		return string(IdentifierSofthashISCC)
	case HeaderZeroxcertZEROXCERTIMPRINT256:
		return string(IdentifierZeroxcertZEROXCERTIMPRINT256)
	case HeaderVarsigNONSTANDARDSIG:
		return string(IdentifierVarsigNONSTANDARDSIG)
	case HeaderVarsigES256K:
		return string(IdentifierVarsigES256K)
	case HeaderVarsigBLS12381G1SIG:
		return string(IdentifierVarsigBLS12381G1SIG)
	case HeaderVarsigBLS12381G2SIG:
		return string(IdentifierVarsigBLS12381G2SIG)
	case HeaderVarsigEDDSA:
		return string(IdentifierVarsigEDDSA)
	case HeaderVarsigEIP191:
		return string(IdentifierVarsigEIP191)
	case HeaderKeyJWKJCSPUB:
		return string(IdentifierKeyJWKJCSPUB)
	case HeaderFilecoinFILCOMMITMENTUNSEALED:
		return string(IdentifierFilecoinFILCOMMITMENTUNSEALED)
	case HeaderFilecoinFILCOMMITMENTSEALED:
		return string(IdentifierFilecoinFILCOMMITMENTSEALED)
	case HeaderMultiaddrPLAINTEXTV2:
		return string(IdentifierMultiaddrPLAINTEXTV2)
	case HeaderHolochainHOLOCHAINADRV0:
		return string(IdentifierHolochainHOLOCHAINADRV0)
	case HeaderHolochainHOLOCHAINADRV1:
		return string(IdentifierHolochainHOLOCHAINADRV1)
	case HeaderHolochainHOLOCHAINKEYV0:
		return string(IdentifierHolochainHOLOCHAINKEYV0)
	case HeaderHolochainHOLOCHAINKEYV1:
		return string(IdentifierHolochainHOLOCHAINKEYV1)
	case HeaderHolochainHOLOCHAINSIGV0:
		return string(IdentifierHolochainHOLOCHAINSIGV0)
	case HeaderHolochainHOLOCHAINSIGV1:
		return string(IdentifierHolochainHOLOCHAINSIGV1)
	case HeaderNamespaceSKYNETNS:
		return string(IdentifierNamespaceSKYNETNS)
	case HeaderNamespaceARWEAVENS:
		return string(IdentifierNamespaceARWEAVENS)
	case HeaderNamespaceSUBSPACENS:
		return string(IdentifierNamespaceSUBSPACENS)
	case HeaderNamespaceKUMANDRANS:
		return string(IdentifierNamespaceKUMANDRANS)
	case HeaderVarsigES256:
		return string(IdentifierVarsigES256)
	case HeaderVarsigES284:
		return string(IdentifierVarsigES284)
	case HeaderVarsigES512:
		return string(IdentifierVarsigES512)
	case HeaderVarsigRS256:
		return string(IdentifierVarsigRS256)
	case HeaderMultiaddrSCION:
		return string(IdentifierMultiaddrSCION)
	default:
		return string(IdentifierUNKNOWN)
	}
}
/*
 * Encode identifier.
 */
func (id Identifier) Code() (uint64) {
	switch id {
	case IdentifierMultihashIDENTITY:
		return uint64(HeaderMultihashIDENTITY)
	case IdentifierCidCIDV1:
		return uint64(HeaderCidCIDV1)
	case IdentifierCidCIDV2:
		return uint64(HeaderCidCIDV2)
	case IdentifierCidCIDV3:
		return uint64(HeaderCidCIDV3)
	case IdentifierMultiaddrIP4:
		return uint64(HeaderMultiaddrIP4)
	case IdentifierMultiaddrTCP:
		return uint64(HeaderMultiaddrTCP)
	case IdentifierMultihashSHA1:
		return uint64(HeaderMultihashSHA1)
	case IdentifierMultihashSHA2256:
		return uint64(HeaderMultihashSHA2256)
	case IdentifierMultihashSHA2512:
		return uint64(HeaderMultihashSHA2512)
	case IdentifierMultihashSHA3512:
		return uint64(HeaderMultihashSHA3512)
	case IdentifierMultihashSHA3384:
		return uint64(HeaderMultihashSHA3384)
	case IdentifierMultihashSHA3256:
		return uint64(HeaderMultihashSHA3256)
	case IdentifierMultihashSHA3224:
		return uint64(HeaderMultihashSHA3224)
	case IdentifierMultihashSHAKE128:
		return uint64(HeaderMultihashSHAKE128)
	case IdentifierMultihashSHAKE256:
		return uint64(HeaderMultihashSHAKE256)
	case IdentifierMultihashKECCAK224:
		return uint64(HeaderMultihashKECCAK224)
	case IdentifierMultihashKECCAK256:
		return uint64(HeaderMultihashKECCAK256)
	case IdentifierMultihashKECCAK384:
		return uint64(HeaderMultihashKECCAK384)
	case IdentifierMultihashKECCAK512:
		return uint64(HeaderMultihashKECCAK512)
	case IdentifierMultihashBLAKE3:
		return uint64(HeaderMultihashBLAKE3)
	case IdentifierMultihashSHA2384:
		return uint64(HeaderMultihashSHA2384)
	case IdentifierMultiaddrDCCP:
		return uint64(HeaderMultiaddrDCCP)
	case IdentifierHashMURMUR3X6464:
		return uint64(HeaderHashMURMUR3X6464)
	case IdentifierHashMURMUR332:
		return uint64(HeaderHashMURMUR332)
	case IdentifierMultiaddrIP6:
		return uint64(HeaderMultiaddrIP6)
	case IdentifierMultiaddrIP6ZONE:
		return uint64(HeaderMultiaddrIP6ZONE)
	case IdentifierMultiaddrIPCIDR:
		return uint64(HeaderMultiaddrIPCIDR)
	case IdentifierNamespacePATH:
		return uint64(HeaderNamespacePATH)
	case IdentifierMultiformatMULTICODEC:
		return uint64(HeaderMultiformatMULTICODEC)
	case IdentifierMultiformatMULTIHASH:
		return uint64(HeaderMultiformatMULTIHASH)
	case IdentifierMultiformatMULTIADDR:
		return uint64(HeaderMultiformatMULTIADDR)
	case IdentifierMultiformatMULTIBASE:
		return uint64(HeaderMultiformatMULTIBASE)
	case IdentifierMultiformatVARSIG:
		return uint64(HeaderMultiformatVARSIG)
	case IdentifierMultiaddrDNS:
		return uint64(HeaderMultiaddrDNS)
	case IdentifierMultiaddrDNS4:
		return uint64(HeaderMultiaddrDNS4)
	case IdentifierMultiaddrDNS6:
		return uint64(HeaderMultiaddrDNS6)
	case IdentifierMultiaddrDNSADDR:
		return uint64(HeaderMultiaddrDNSADDR)
	case IdentifierSerializationPROTOBUF:
		return uint64(HeaderSerializationPROTOBUF)
	case IdentifierIpldCBOR:
		return uint64(HeaderIpldCBOR)
	case IdentifierIpldRAW:
		return uint64(HeaderIpldRAW)
	case IdentifierMultihashDBLSHA2256:
		return uint64(HeaderMultihashDBLSHA2256)
	case IdentifierSerializationRLP:
		return uint64(HeaderSerializationRLP)
	case IdentifierSerializationBENCODE:
		return uint64(HeaderSerializationBENCODE)
	case IdentifierIpldDAGPB:
		return uint64(HeaderIpldDAGPB)
	case IdentifierIpldDAGCBOR:
		return uint64(HeaderIpldDAGCBOR)
	case IdentifierIpldLIBP2PKEY:
		return uint64(HeaderIpldLIBP2PKEY)
	case IdentifierIpldGITRAW:
		return uint64(HeaderIpldGITRAW)
	case IdentifierIpldTORRENTINFO:
		return uint64(HeaderIpldTORRENTINFO)
	case IdentifierIpldTORRENTFILE:
		return uint64(HeaderIpldTORRENTFILE)
	case IdentifierIpldLEOFCOINBLOCK:
		return uint64(HeaderIpldLEOFCOINBLOCK)
	case IdentifierIpldLEOFCOINTX:
		return uint64(HeaderIpldLEOFCOINTX)
	case IdentifierIpldLEOFCOINPR:
		return uint64(HeaderIpldLEOFCOINPR)
	case IdentifierMultiaddrSCTP:
		return uint64(HeaderMultiaddrSCTP)
	case IdentifierIpldDAGJOSE:
		return uint64(HeaderIpldDAGJOSE)
	case IdentifierIpldDAGCOSE:
		return uint64(HeaderIpldDAGCOSE)
	case IdentifierNamespaceLBRY:
		return uint64(HeaderNamespaceLBRY)
	case IdentifierIpldETHBLOCK:
		return uint64(HeaderIpldETHBLOCK)
	case IdentifierIpldETHBLOCKLIST:
		return uint64(HeaderIpldETHBLOCKLIST)
	case IdentifierIpldETHTXTRIE:
		return uint64(HeaderIpldETHTXTRIE)
	case IdentifierIpldETHTX:
		return uint64(HeaderIpldETHTX)
	case IdentifierIpldETHTXRECEIPTTRIE:
		return uint64(HeaderIpldETHTXRECEIPTTRIE)
	case IdentifierIpldETHTXRECEIPT:
		return uint64(HeaderIpldETHTXRECEIPT)
	case IdentifierIpldETHSTATETRIE:
		return uint64(HeaderIpldETHSTATETRIE)
	case IdentifierIpldETHACCOUNTSNAPSHOT:
		return uint64(HeaderIpldETHACCOUNTSNAPSHOT)
	case IdentifierIpldETHSTORAGETRIE:
		return uint64(HeaderIpldETHSTORAGETRIE)
	case IdentifierIpldETHRECEIPTLOGTRIE:
		return uint64(HeaderIpldETHRECEIPTLOGTRIE)
	case IdentifierIpldETHRECEIPTLOG:
		return uint64(HeaderIpldETHRECEIPTLOG)
	case IdentifierKeyAES128:
		return uint64(HeaderKeyAES128)
	case IdentifierKeyAES192:
		return uint64(HeaderKeyAES192)
	case IdentifierKeyAES256:
		return uint64(HeaderKeyAES256)
	case IdentifierKeyCHACHA128:
		return uint64(HeaderKeyCHACHA128)
	case IdentifierKeyCHACHA256:
		return uint64(HeaderKeyCHACHA256)
	case IdentifierIpldBITCOINBLOCK:
		return uint64(HeaderIpldBITCOINBLOCK)
	case IdentifierIpldBITCOINTX:
		return uint64(HeaderIpldBITCOINTX)
	case IdentifierIpldBITCOINWITNESSCOMMITMENT:
		return uint64(HeaderIpldBITCOINWITNESSCOMMITMENT)
	case IdentifierIpldZCASHBLOCK:
		return uint64(HeaderIpldZCASHBLOCK)
	case IdentifierIpldZCASHTX:
		return uint64(HeaderIpldZCASHTX)
	case IdentifierMultiformatCAIP50:
		return uint64(HeaderMultiformatCAIP50)
	case IdentifierNamespaceSTREAMID:
		return uint64(HeaderNamespaceSTREAMID)
	case IdentifierIpldSTELLARBLOCK:
		return uint64(HeaderIpldSTELLARBLOCK)
	case IdentifierIpldSTELLARTX:
		return uint64(HeaderIpldSTELLARTX)
	case IdentifierMultihashMD4:
		return uint64(HeaderMultihashMD4)
	case IdentifierMultihashMD5:
		return uint64(HeaderMultihashMD5)
	case IdentifierIpldDECREDBLOCK:
		return uint64(HeaderIpldDECREDBLOCK)
	case IdentifierIpldDECREDTX:
		return uint64(HeaderIpldDECREDTX)
	case IdentifierNamespaceIPLD:
		return uint64(HeaderNamespaceIPLD)
	case IdentifierNamespaceIPFS:
		return uint64(HeaderNamespaceIPFS)
	case IdentifierNamespaceSWARM:
		return uint64(HeaderNamespaceSWARM)
	case IdentifierNamespaceIPNS:
		return uint64(HeaderNamespaceIPNS)
	case IdentifierNamespaceZERONET:
		return uint64(HeaderNamespaceZERONET)
	case IdentifierKeySECP256K1PUB:
		return uint64(HeaderKeySECP256K1PUB)
	case IdentifierNamespaceDNSLINK:
		return uint64(HeaderNamespaceDNSLINK)
	case IdentifierKeyBLS12381G1PUB:
		return uint64(HeaderKeyBLS12381G1PUB)
	case IdentifierKeyBLS12381G2PUB:
		return uint64(HeaderKeyBLS12381G2PUB)
	case IdentifierKeyX25519PUB:
		return uint64(HeaderKeyX25519PUB)
	case IdentifierKeyED25519PUB:
		return uint64(HeaderKeyED25519PUB)
	case IdentifierKeyBLS12381G1G2PUB:
		return uint64(HeaderKeyBLS12381G1G2PUB)
	case IdentifierKeySR25519PUB:
		return uint64(HeaderKeySR25519PUB)
	case IdentifierIpldDASHBLOCK:
		return uint64(HeaderIpldDASHBLOCK)
	case IdentifierIpldDASHTX:
		return uint64(HeaderIpldDASHTX)
	case IdentifierIpldSWARMMANIFEST:
		return uint64(HeaderIpldSWARMMANIFEST)
	case IdentifierIpldSWARMFEED:
		return uint64(HeaderIpldSWARMFEED)
	case IdentifierIpldBEESON:
		return uint64(HeaderIpldBEESON)
	case IdentifierMultiaddrUDP:
		return uint64(HeaderMultiaddrUDP)
	case IdentifierMultiaddrP2PWEBRTCSTAR:
		return uint64(HeaderMultiaddrP2PWEBRTCSTAR)
	case IdentifierMultiaddrP2PWEBRTCDIRECT:
		return uint64(HeaderMultiaddrP2PWEBRTCDIRECT)
	case IdentifierMultiaddrP2PSTARDUST:
		return uint64(HeaderMultiaddrP2PSTARDUST)
	case IdentifierMultiaddrWEBRTCDIRECT:
		return uint64(HeaderMultiaddrWEBRTCDIRECT)
	case IdentifierMultiaddrWEBRTC:
		return uint64(HeaderMultiaddrWEBRTC)
	case IdentifierMultiaddrP2PCIRCUIT:
		return uint64(HeaderMultiaddrP2PCIRCUIT)
	case IdentifierIpldDAGJSON:
		return uint64(HeaderIpldDAGJSON)
	case IdentifierMultiaddrUDT:
		return uint64(HeaderMultiaddrUDT)
	case IdentifierMultiaddrUTP:
		return uint64(HeaderMultiaddrUTP)
	case IdentifierHashCRC32:
		return uint64(HeaderHashCRC32)
	case IdentifierHashCRC64ECMA:
		return uint64(HeaderHashCRC64ECMA)
	case IdentifierMultiaddrUNIX:
		return uint64(HeaderMultiaddrUNIX)
	case IdentifierMultiaddrTHREAD:
		return uint64(HeaderMultiaddrTHREAD)
	case IdentifierMultiaddrP2P:
		return uint64(HeaderMultiaddrP2P)
	case IdentifierMultiaddrHTTPS:
		return uint64(HeaderMultiaddrHTTPS)
	case IdentifierMultiaddrONION:
		return uint64(HeaderMultiaddrONION)
	case IdentifierMultiaddrONION3:
		return uint64(HeaderMultiaddrONION3)
	case IdentifierMultiaddrGARLIC64:
		return uint64(HeaderMultiaddrGARLIC64)
	case IdentifierMultiaddrGARLIC32:
		return uint64(HeaderMultiaddrGARLIC32)
	case IdentifierMultiaddrTLS:
		return uint64(HeaderMultiaddrTLS)
	case IdentifierMultiaddrSNI:
		return uint64(HeaderMultiaddrSNI)
	case IdentifierMultiaddrNOISE:
		return uint64(HeaderMultiaddrNOISE)
	case IdentifierMultiaddrSHS:
		return uint64(HeaderMultiaddrSHS)
	case IdentifierMultiaddrQUIC:
		return uint64(HeaderMultiaddrQUIC)
	case IdentifierMultiaddrQUICV1:
		return uint64(HeaderMultiaddrQUICV1)
	case IdentifierMultiaddrWEBTRANSPORT:
		return uint64(HeaderMultiaddrWEBTRANSPORT)
	case IdentifierMultiaddrCERTHASH:
		return uint64(HeaderMultiaddrCERTHASH)
	case IdentifierMultiaddrWS:
		return uint64(HeaderMultiaddrWS)
	case IdentifierMultiaddrWSS:
		return uint64(HeaderMultiaddrWSS)
	case IdentifierMultiaddrP2PWEBSOCKETSTAR:
		return uint64(HeaderMultiaddrP2PWEBSOCKETSTAR)
	case IdentifierMultiaddrHTTP:
		return uint64(HeaderMultiaddrHTTP)
	case IdentifierIpldSWHID1SNP:
		return uint64(HeaderIpldSWHID1SNP)
	case IdentifierIpldJSON:
		return uint64(HeaderIpldJSON)
	case IdentifierSerializationMESSAGEPACK:
		return uint64(HeaderSerializationMESSAGEPACK)
	case IdentifierSerializationCAR:
		return uint64(HeaderSerializationCAR)
	case IdentifierSerializationIPNSRECORD:
		return uint64(HeaderSerializationIPNSRECORD)
	case IdentifierLibp2pLIBP2PPEERRECORD:
		return uint64(HeaderLibp2pLIBP2PPEERRECORD)
	case IdentifierLibp2pLIBP2PRELAYRSVP:
		return uint64(HeaderLibp2pLIBP2PRELAYRSVP)
	case IdentifierLibp2pMEMORYTRANSPORT:
		return uint64(HeaderLibp2pMEMORYTRANSPORT)
	case IdentifierSerializationCARINDEXSORTED:
		return uint64(HeaderSerializationCARINDEXSORTED)
	case IdentifierSerializationCARMULTIHASHINDEXSORTED:
		return uint64(HeaderSerializationCARMULTIHASHINDEXSORTED)
	case IdentifierTransportTRANSPORTBITSWAP:
		return uint64(HeaderTransportTRANSPORTBITSWAP)
	case IdentifierTransportTRANSPORTGRAPHSYNCFILECOINV1:
		return uint64(HeaderTransportTRANSPORTGRAPHSYNCFILECOINV1)
	case IdentifierTransportTRANSPORTIPFSGATEWAYHTTP:
		return uint64(HeaderTransportTRANSPORTIPFSGATEWAYHTTP)
	case IdentifierMultiformatMULTIDID:
		return uint64(HeaderMultiformatMULTIDID)
	case IdentifierMultihashSHA2256TRUNC254PADDED:
		return uint64(HeaderMultihashSHA2256TRUNC254PADDED)
	case IdentifierMultihashSHA2224:
		return uint64(HeaderMultihashSHA2224)
	case IdentifierMultihashSHA2512224:
		return uint64(HeaderMultihashSHA2512224)
	case IdentifierMultihashSHA2512256:
		return uint64(HeaderMultihashSHA2512256)
	case IdentifierHashMURMUR3X64128:
		return uint64(HeaderHashMURMUR3X64128)
	case IdentifierMultihashRIPEMD128:
		return uint64(HeaderMultihashRIPEMD128)
	case IdentifierMultihashRIPEMD160:
		return uint64(HeaderMultihashRIPEMD160)
	case IdentifierMultihashRIPEMD256:
		return uint64(HeaderMultihashRIPEMD256)
	case IdentifierMultihashRIPEMD320:
		return uint64(HeaderMultihashRIPEMD320)
	case IdentifierMultihashX11:
		return uint64(HeaderMultihashX11)
	case IdentifierKeyP256PUB:
		return uint64(HeaderKeyP256PUB)
	case IdentifierKeyP384PUB:
		return uint64(HeaderKeyP384PUB)
	case IdentifierKeyP521PUB:
		return uint64(HeaderKeyP521PUB)
	case IdentifierKeyED448PUB:
		return uint64(HeaderKeyED448PUB)
	case IdentifierKeyX448PUB:
		return uint64(HeaderKeyX448PUB)
	case IdentifierKeyRSAPUB:
		return uint64(HeaderKeyRSAPUB)
	case IdentifierKeySM2PUB:
		return uint64(HeaderKeySM2PUB)
	case IdentifierKeyED25519PRIV:
		return uint64(HeaderKeyED25519PRIV)
	case IdentifierKeySECP256K1PRIV:
		return uint64(HeaderKeySECP256K1PRIV)
	case IdentifierKeyX25519PRIV:
		return uint64(HeaderKeyX25519PRIV)
	case IdentifierKeySR25519PRIV:
		return uint64(HeaderKeySR25519PRIV)
	case IdentifierKeyRSAPRIV:
		return uint64(HeaderKeyRSAPRIV)
	case IdentifierKeyP256PRIV:
		return uint64(HeaderKeyP256PRIV)
	case IdentifierKeyP384PRIV:
		return uint64(HeaderKeyP384PRIV)
	case IdentifierKeyP521PRIV:
		return uint64(HeaderKeyP521PRIV)
	case IdentifierMultihashKANGAROOTWELVE:
		return uint64(HeaderMultihashKANGAROOTWELVE)
	case IdentifierEncryptionAESGCM256:
		return uint64(HeaderEncryptionAESGCM256)
	case IdentifierMultiaddrSILVERPINE:
		return uint64(HeaderMultiaddrSILVERPINE)
	case IdentifierMultihashSM3256:
		return uint64(HeaderMultihashSM3256)
	case IdentifierHashSHA256A:
		return uint64(HeaderHashSHA256A)
	case IdentifierMultihashBLAKE2B8:
		return uint64(HeaderMultihashBLAKE2B8)
	case IdentifierMultihashBLAKE2B16:
		return uint64(HeaderMultihashBLAKE2B16)
	case IdentifierMultihashBLAKE2B24:
		return uint64(HeaderMultihashBLAKE2B24)
	case IdentifierMultihashBLAKE2B32:
		return uint64(HeaderMultihashBLAKE2B32)
	case IdentifierMultihashBLAKE2B40:
		return uint64(HeaderMultihashBLAKE2B40)
	case IdentifierMultihashBLAKE2B48:
		return uint64(HeaderMultihashBLAKE2B48)
	case IdentifierMultihashBLAKE2B56:
		return uint64(HeaderMultihashBLAKE2B56)
	case IdentifierMultihashBLAKE2B64:
		return uint64(HeaderMultihashBLAKE2B64)
	case IdentifierMultihashBLAKE2B72:
		return uint64(HeaderMultihashBLAKE2B72)
	case IdentifierMultihashBLAKE2B80:
		return uint64(HeaderMultihashBLAKE2B80)
	case IdentifierMultihashBLAKE2B88:
		return uint64(HeaderMultihashBLAKE2B88)
	case IdentifierMultihashBLAKE2B96:
		return uint64(HeaderMultihashBLAKE2B96)
	case IdentifierMultihashBLAKE2B104:
		return uint64(HeaderMultihashBLAKE2B104)
	case IdentifierMultihashBLAKE2B112:
		return uint64(HeaderMultihashBLAKE2B112)
	case IdentifierMultihashBLAKE2B120:
		return uint64(HeaderMultihashBLAKE2B120)
	case IdentifierMultihashBLAKE2B128:
		return uint64(HeaderMultihashBLAKE2B128)
	case IdentifierMultihashBLAKE2B136:
		return uint64(HeaderMultihashBLAKE2B136)
	case IdentifierMultihashBLAKE2B144:
		return uint64(HeaderMultihashBLAKE2B144)
	case IdentifierMultihashBLAKE2B152:
		return uint64(HeaderMultihashBLAKE2B152)
	case IdentifierMultihashBLAKE2B160:
		return uint64(HeaderMultihashBLAKE2B160)
	case IdentifierMultihashBLAKE2B168:
		return uint64(HeaderMultihashBLAKE2B168)
	case IdentifierMultihashBLAKE2B176:
		return uint64(HeaderMultihashBLAKE2B176)
	case IdentifierMultihashBLAKE2B184:
		return uint64(HeaderMultihashBLAKE2B184)
	case IdentifierMultihashBLAKE2B192:
		return uint64(HeaderMultihashBLAKE2B192)
	case IdentifierMultihashBLAKE2B200:
		return uint64(HeaderMultihashBLAKE2B200)
	case IdentifierMultihashBLAKE2B208:
		return uint64(HeaderMultihashBLAKE2B208)
	case IdentifierMultihashBLAKE2B216:
		return uint64(HeaderMultihashBLAKE2B216)
	case IdentifierMultihashBLAKE2B224:
		return uint64(HeaderMultihashBLAKE2B224)
	case IdentifierMultihashBLAKE2B232:
		return uint64(HeaderMultihashBLAKE2B232)
	case IdentifierMultihashBLAKE2B240:
		return uint64(HeaderMultihashBLAKE2B240)
	case IdentifierMultihashBLAKE2B248:
		return uint64(HeaderMultihashBLAKE2B248)
	case IdentifierMultihashBLAKE2B256:
		return uint64(HeaderMultihashBLAKE2B256)
	case IdentifierMultihashBLAKE2B264:
		return uint64(HeaderMultihashBLAKE2B264)
	case IdentifierMultihashBLAKE2B272:
		return uint64(HeaderMultihashBLAKE2B272)
	case IdentifierMultihashBLAKE2B280:
		return uint64(HeaderMultihashBLAKE2B280)
	case IdentifierMultihashBLAKE2B288:
		return uint64(HeaderMultihashBLAKE2B288)
	case IdentifierMultihashBLAKE2B296:
		return uint64(HeaderMultihashBLAKE2B296)
	case IdentifierMultihashBLAKE2B304:
		return uint64(HeaderMultihashBLAKE2B304)
	case IdentifierMultihashBLAKE2B312:
		return uint64(HeaderMultihashBLAKE2B312)
	case IdentifierMultihashBLAKE2B320:
		return uint64(HeaderMultihashBLAKE2B320)
	case IdentifierMultihashBLAKE2B328:
		return uint64(HeaderMultihashBLAKE2B328)
	case IdentifierMultihashBLAKE2B336:
		return uint64(HeaderMultihashBLAKE2B336)
	case IdentifierMultihashBLAKE2B344:
		return uint64(HeaderMultihashBLAKE2B344)
	case IdentifierMultihashBLAKE2B352:
		return uint64(HeaderMultihashBLAKE2B352)
	case IdentifierMultihashBLAKE2B360:
		return uint64(HeaderMultihashBLAKE2B360)
	case IdentifierMultihashBLAKE2B368:
		return uint64(HeaderMultihashBLAKE2B368)
	case IdentifierMultihashBLAKE2B376:
		return uint64(HeaderMultihashBLAKE2B376)
	case IdentifierMultihashBLAKE2B384:
		return uint64(HeaderMultihashBLAKE2B384)
	case IdentifierMultihashBLAKE2B392:
		return uint64(HeaderMultihashBLAKE2B392)
	case IdentifierMultihashBLAKE2B400:
		return uint64(HeaderMultihashBLAKE2B400)
	case IdentifierMultihashBLAKE2B408:
		return uint64(HeaderMultihashBLAKE2B408)
	case IdentifierMultihashBLAKE2B416:
		return uint64(HeaderMultihashBLAKE2B416)
	case IdentifierMultihashBLAKE2B424:
		return uint64(HeaderMultihashBLAKE2B424)
	case IdentifierMultihashBLAKE2B432:
		return uint64(HeaderMultihashBLAKE2B432)
	case IdentifierMultihashBLAKE2B440:
		return uint64(HeaderMultihashBLAKE2B440)
	case IdentifierMultihashBLAKE2B448:
		return uint64(HeaderMultihashBLAKE2B448)
	case IdentifierMultihashBLAKE2B456:
		return uint64(HeaderMultihashBLAKE2B456)
	case IdentifierMultihashBLAKE2B464:
		return uint64(HeaderMultihashBLAKE2B464)
	case IdentifierMultihashBLAKE2B472:
		return uint64(HeaderMultihashBLAKE2B472)
	case IdentifierMultihashBLAKE2B480:
		return uint64(HeaderMultihashBLAKE2B480)
	case IdentifierMultihashBLAKE2B488:
		return uint64(HeaderMultihashBLAKE2B488)
	case IdentifierMultihashBLAKE2B496:
		return uint64(HeaderMultihashBLAKE2B496)
	case IdentifierMultihashBLAKE2B504:
		return uint64(HeaderMultihashBLAKE2B504)
	case IdentifierMultihashBLAKE2B512:
		return uint64(HeaderMultihashBLAKE2B512)
	case IdentifierMultihashBLAKE2S8:
		return uint64(HeaderMultihashBLAKE2S8)
	case IdentifierMultihashBLAKE2S16:
		return uint64(HeaderMultihashBLAKE2S16)
	case IdentifierMultihashBLAKE2S24:
		return uint64(HeaderMultihashBLAKE2S24)
	case IdentifierMultihashBLAKE2S32:
		return uint64(HeaderMultihashBLAKE2S32)
	case IdentifierMultihashBLAKE2S40:
		return uint64(HeaderMultihashBLAKE2S40)
	case IdentifierMultihashBLAKE2S48:
		return uint64(HeaderMultihashBLAKE2S48)
	case IdentifierMultihashBLAKE2S56:
		return uint64(HeaderMultihashBLAKE2S56)
	case IdentifierMultihashBLAKE2S64:
		return uint64(HeaderMultihashBLAKE2S64)
	case IdentifierMultihashBLAKE2S72:
		return uint64(HeaderMultihashBLAKE2S72)
	case IdentifierMultihashBLAKE2S80:
		return uint64(HeaderMultihashBLAKE2S80)
	case IdentifierMultihashBLAKE2S88:
		return uint64(HeaderMultihashBLAKE2S88)
	case IdentifierMultihashBLAKE2S96:
		return uint64(HeaderMultihashBLAKE2S96)
	case IdentifierMultihashBLAKE2S104:
		return uint64(HeaderMultihashBLAKE2S104)
	case IdentifierMultihashBLAKE2S112:
		return uint64(HeaderMultihashBLAKE2S112)
	case IdentifierMultihashBLAKE2S120:
		return uint64(HeaderMultihashBLAKE2S120)
	case IdentifierMultihashBLAKE2S128:
		return uint64(HeaderMultihashBLAKE2S128)
	case IdentifierMultihashBLAKE2S136:
		return uint64(HeaderMultihashBLAKE2S136)
	case IdentifierMultihashBLAKE2S144:
		return uint64(HeaderMultihashBLAKE2S144)
	case IdentifierMultihashBLAKE2S152:
		return uint64(HeaderMultihashBLAKE2S152)
	case IdentifierMultihashBLAKE2S160:
		return uint64(HeaderMultihashBLAKE2S160)
	case IdentifierMultihashBLAKE2S168:
		return uint64(HeaderMultihashBLAKE2S168)
	case IdentifierMultihashBLAKE2S176:
		return uint64(HeaderMultihashBLAKE2S176)
	case IdentifierMultihashBLAKE2S184:
		return uint64(HeaderMultihashBLAKE2S184)
	case IdentifierMultihashBLAKE2S192:
		return uint64(HeaderMultihashBLAKE2S192)
	case IdentifierMultihashBLAKE2S200:
		return uint64(HeaderMultihashBLAKE2S200)
	case IdentifierMultihashBLAKE2S208:
		return uint64(HeaderMultihashBLAKE2S208)
	case IdentifierMultihashBLAKE2S216:
		return uint64(HeaderMultihashBLAKE2S216)
	case IdentifierMultihashBLAKE2S224:
		return uint64(HeaderMultihashBLAKE2S224)
	case IdentifierMultihashBLAKE2S232:
		return uint64(HeaderMultihashBLAKE2S232)
	case IdentifierMultihashBLAKE2S240:
		return uint64(HeaderMultihashBLAKE2S240)
	case IdentifierMultihashBLAKE2S248:
		return uint64(HeaderMultihashBLAKE2S248)
	case IdentifierMultihashBLAKE2S256:
		return uint64(HeaderMultihashBLAKE2S256)
	case IdentifierMultihashSKEIN2568:
		return uint64(HeaderMultihashSKEIN2568)
	case IdentifierMultihashSKEIN25616:
		return uint64(HeaderMultihashSKEIN25616)
	case IdentifierMultihashSKEIN25624:
		return uint64(HeaderMultihashSKEIN25624)
	case IdentifierMultihashSKEIN25632:
		return uint64(HeaderMultihashSKEIN25632)
	case IdentifierMultihashSKEIN25640:
		return uint64(HeaderMultihashSKEIN25640)
	case IdentifierMultihashSKEIN25648:
		return uint64(HeaderMultihashSKEIN25648)
	case IdentifierMultihashSKEIN25656:
		return uint64(HeaderMultihashSKEIN25656)
	case IdentifierMultihashSKEIN25664:
		return uint64(HeaderMultihashSKEIN25664)
	case IdentifierMultihashSKEIN25672:
		return uint64(HeaderMultihashSKEIN25672)
	case IdentifierMultihashSKEIN25680:
		return uint64(HeaderMultihashSKEIN25680)
	case IdentifierMultihashSKEIN25688:
		return uint64(HeaderMultihashSKEIN25688)
	case IdentifierMultihashSKEIN25696:
		return uint64(HeaderMultihashSKEIN25696)
	case IdentifierMultihashSKEIN256104:
		return uint64(HeaderMultihashSKEIN256104)
	case IdentifierMultihashSKEIN256112:
		return uint64(HeaderMultihashSKEIN256112)
	case IdentifierMultihashSKEIN256120:
		return uint64(HeaderMultihashSKEIN256120)
	case IdentifierMultihashSKEIN256128:
		return uint64(HeaderMultihashSKEIN256128)
	case IdentifierMultihashSKEIN256136:
		return uint64(HeaderMultihashSKEIN256136)
	case IdentifierMultihashSKEIN256144:
		return uint64(HeaderMultihashSKEIN256144)
	case IdentifierMultihashSKEIN256152:
		return uint64(HeaderMultihashSKEIN256152)
	case IdentifierMultihashSKEIN256160:
		return uint64(HeaderMultihashSKEIN256160)
	case IdentifierMultihashSKEIN256168:
		return uint64(HeaderMultihashSKEIN256168)
	case IdentifierMultihashSKEIN256176:
		return uint64(HeaderMultihashSKEIN256176)
	case IdentifierMultihashSKEIN256184:
		return uint64(HeaderMultihashSKEIN256184)
	case IdentifierMultihashSKEIN256192:
		return uint64(HeaderMultihashSKEIN256192)
	case IdentifierMultihashSKEIN256200:
		return uint64(HeaderMultihashSKEIN256200)
	case IdentifierMultihashSKEIN256208:
		return uint64(HeaderMultihashSKEIN256208)
	case IdentifierMultihashSKEIN256216:
		return uint64(HeaderMultihashSKEIN256216)
	case IdentifierMultihashSKEIN256224:
		return uint64(HeaderMultihashSKEIN256224)
	case IdentifierMultihashSKEIN256232:
		return uint64(HeaderMultihashSKEIN256232)
	case IdentifierMultihashSKEIN256240:
		return uint64(HeaderMultihashSKEIN256240)
	case IdentifierMultihashSKEIN256248:
		return uint64(HeaderMultihashSKEIN256248)
	case IdentifierMultihashSKEIN256256:
		return uint64(HeaderMultihashSKEIN256256)
	case IdentifierMultihashSKEIN5128:
		return uint64(HeaderMultihashSKEIN5128)
	case IdentifierMultihashSKEIN51216:
		return uint64(HeaderMultihashSKEIN51216)
	case IdentifierMultihashSKEIN51224:
		return uint64(HeaderMultihashSKEIN51224)
	case IdentifierMultihashSKEIN51232:
		return uint64(HeaderMultihashSKEIN51232)
	case IdentifierMultihashSKEIN51240:
		return uint64(HeaderMultihashSKEIN51240)
	case IdentifierMultihashSKEIN51248:
		return uint64(HeaderMultihashSKEIN51248)
	case IdentifierMultihashSKEIN51256:
		return uint64(HeaderMultihashSKEIN51256)
	case IdentifierMultihashSKEIN51264:
		return uint64(HeaderMultihashSKEIN51264)
	case IdentifierMultihashSKEIN51272:
		return uint64(HeaderMultihashSKEIN51272)
	case IdentifierMultihashSKEIN51280:
		return uint64(HeaderMultihashSKEIN51280)
	case IdentifierMultihashSKEIN51288:
		return uint64(HeaderMultihashSKEIN51288)
	case IdentifierMultihashSKEIN51296:
		return uint64(HeaderMultihashSKEIN51296)
	case IdentifierMultihashSKEIN512104:
		return uint64(HeaderMultihashSKEIN512104)
	case IdentifierMultihashSKEIN512112:
		return uint64(HeaderMultihashSKEIN512112)
	case IdentifierMultihashSKEIN512120:
		return uint64(HeaderMultihashSKEIN512120)
	case IdentifierMultihashSKEIN512128:
		return uint64(HeaderMultihashSKEIN512128)
	case IdentifierMultihashSKEIN512136:
		return uint64(HeaderMultihashSKEIN512136)
	case IdentifierMultihashSKEIN512144:
		return uint64(HeaderMultihashSKEIN512144)
	case IdentifierMultihashSKEIN512152:
		return uint64(HeaderMultihashSKEIN512152)
	case IdentifierMultihashSKEIN512160:
		return uint64(HeaderMultihashSKEIN512160)
	case IdentifierMultihashSKEIN512168:
		return uint64(HeaderMultihashSKEIN512168)
	case IdentifierMultihashSKEIN512176:
		return uint64(HeaderMultihashSKEIN512176)
	case IdentifierMultihashSKEIN512184:
		return uint64(HeaderMultihashSKEIN512184)
	case IdentifierMultihashSKEIN512192:
		return uint64(HeaderMultihashSKEIN512192)
	case IdentifierMultihashSKEIN512200:
		return uint64(HeaderMultihashSKEIN512200)
	case IdentifierMultihashSKEIN512208:
		return uint64(HeaderMultihashSKEIN512208)
	case IdentifierMultihashSKEIN512216:
		return uint64(HeaderMultihashSKEIN512216)
	case IdentifierMultihashSKEIN512224:
		return uint64(HeaderMultihashSKEIN512224)
	case IdentifierMultihashSKEIN512232:
		return uint64(HeaderMultihashSKEIN512232)
	case IdentifierMultihashSKEIN512240:
		return uint64(HeaderMultihashSKEIN512240)
	case IdentifierMultihashSKEIN512248:
		return uint64(HeaderMultihashSKEIN512248)
	case IdentifierMultihashSKEIN512256:
		return uint64(HeaderMultihashSKEIN512256)
	case IdentifierMultihashSKEIN512264:
		return uint64(HeaderMultihashSKEIN512264)
	case IdentifierMultihashSKEIN512272:
		return uint64(HeaderMultihashSKEIN512272)
	case IdentifierMultihashSKEIN512280:
		return uint64(HeaderMultihashSKEIN512280)
	case IdentifierMultihashSKEIN512288:
		return uint64(HeaderMultihashSKEIN512288)
	case IdentifierMultihashSKEIN512296:
		return uint64(HeaderMultihashSKEIN512296)
	case IdentifierMultihashSKEIN512304:
		return uint64(HeaderMultihashSKEIN512304)
	case IdentifierMultihashSKEIN512312:
		return uint64(HeaderMultihashSKEIN512312)
	case IdentifierMultihashSKEIN512320:
		return uint64(HeaderMultihashSKEIN512320)
	case IdentifierMultihashSKEIN512328:
		return uint64(HeaderMultihashSKEIN512328)
	case IdentifierMultihashSKEIN512336:
		return uint64(HeaderMultihashSKEIN512336)
	case IdentifierMultihashSKEIN512344:
		return uint64(HeaderMultihashSKEIN512344)
	case IdentifierMultihashSKEIN512352:
		return uint64(HeaderMultihashSKEIN512352)
	case IdentifierMultihashSKEIN512360:
		return uint64(HeaderMultihashSKEIN512360)
	case IdentifierMultihashSKEIN512368:
		return uint64(HeaderMultihashSKEIN512368)
	case IdentifierMultihashSKEIN512376:
		return uint64(HeaderMultihashSKEIN512376)
	case IdentifierMultihashSKEIN512384:
		return uint64(HeaderMultihashSKEIN512384)
	case IdentifierMultihashSKEIN512392:
		return uint64(HeaderMultihashSKEIN512392)
	case IdentifierMultihashSKEIN512400:
		return uint64(HeaderMultihashSKEIN512400)
	case IdentifierMultihashSKEIN512408:
		return uint64(HeaderMultihashSKEIN512408)
	case IdentifierMultihashSKEIN512416:
		return uint64(HeaderMultihashSKEIN512416)
	case IdentifierMultihashSKEIN512424:
		return uint64(HeaderMultihashSKEIN512424)
	case IdentifierMultihashSKEIN512432:
		return uint64(HeaderMultihashSKEIN512432)
	case IdentifierMultihashSKEIN512440:
		return uint64(HeaderMultihashSKEIN512440)
	case IdentifierMultihashSKEIN512448:
		return uint64(HeaderMultihashSKEIN512448)
	case IdentifierMultihashSKEIN512456:
		return uint64(HeaderMultihashSKEIN512456)
	case IdentifierMultihashSKEIN512464:
		return uint64(HeaderMultihashSKEIN512464)
	case IdentifierMultihashSKEIN512472:
		return uint64(HeaderMultihashSKEIN512472)
	case IdentifierMultihashSKEIN512480:
		return uint64(HeaderMultihashSKEIN512480)
	case IdentifierMultihashSKEIN512488:
		return uint64(HeaderMultihashSKEIN512488)
	case IdentifierMultihashSKEIN512496:
		return uint64(HeaderMultihashSKEIN512496)
	case IdentifierMultihashSKEIN512504:
		return uint64(HeaderMultihashSKEIN512504)
	case IdentifierMultihashSKEIN512512:
		return uint64(HeaderMultihashSKEIN512512)
	case IdentifierMultihashSKEIN10248:
		return uint64(HeaderMultihashSKEIN10248)
	case IdentifierMultihashSKEIN102416:
		return uint64(HeaderMultihashSKEIN102416)
	case IdentifierMultihashSKEIN102424:
		return uint64(HeaderMultihashSKEIN102424)
	case IdentifierMultihashSKEIN102432:
		return uint64(HeaderMultihashSKEIN102432)
	case IdentifierMultihashSKEIN102440:
		return uint64(HeaderMultihashSKEIN102440)
	case IdentifierMultihashSKEIN102448:
		return uint64(HeaderMultihashSKEIN102448)
	case IdentifierMultihashSKEIN102456:
		return uint64(HeaderMultihashSKEIN102456)
	case IdentifierMultihashSKEIN102464:
		return uint64(HeaderMultihashSKEIN102464)
	case IdentifierMultihashSKEIN102472:
		return uint64(HeaderMultihashSKEIN102472)
	case IdentifierMultihashSKEIN102480:
		return uint64(HeaderMultihashSKEIN102480)
	case IdentifierMultihashSKEIN102488:
		return uint64(HeaderMultihashSKEIN102488)
	case IdentifierMultihashSKEIN102496:
		return uint64(HeaderMultihashSKEIN102496)
	case IdentifierMultihashSKEIN1024104:
		return uint64(HeaderMultihashSKEIN1024104)
	case IdentifierMultihashSKEIN1024112:
		return uint64(HeaderMultihashSKEIN1024112)
	case IdentifierMultihashSKEIN1024120:
		return uint64(HeaderMultihashSKEIN1024120)
	case IdentifierMultihashSKEIN1024128:
		return uint64(HeaderMultihashSKEIN1024128)
	case IdentifierMultihashSKEIN1024136:
		return uint64(HeaderMultihashSKEIN1024136)
	case IdentifierMultihashSKEIN1024144:
		return uint64(HeaderMultihashSKEIN1024144)
	case IdentifierMultihashSKEIN1024152:
		return uint64(HeaderMultihashSKEIN1024152)
	case IdentifierMultihashSKEIN1024160:
		return uint64(HeaderMultihashSKEIN1024160)
	case IdentifierMultihashSKEIN1024168:
		return uint64(HeaderMultihashSKEIN1024168)
	case IdentifierMultihashSKEIN1024176:
		return uint64(HeaderMultihashSKEIN1024176)
	case IdentifierMultihashSKEIN1024184:
		return uint64(HeaderMultihashSKEIN1024184)
	case IdentifierMultihashSKEIN1024192:
		return uint64(HeaderMultihashSKEIN1024192)
	case IdentifierMultihashSKEIN1024200:
		return uint64(HeaderMultihashSKEIN1024200)
	case IdentifierMultihashSKEIN1024208:
		return uint64(HeaderMultihashSKEIN1024208)
	case IdentifierMultihashSKEIN1024216:
		return uint64(HeaderMultihashSKEIN1024216)
	case IdentifierMultihashSKEIN1024224:
		return uint64(HeaderMultihashSKEIN1024224)
	case IdentifierMultihashSKEIN1024232:
		return uint64(HeaderMultihashSKEIN1024232)
	case IdentifierMultihashSKEIN1024240:
		return uint64(HeaderMultihashSKEIN1024240)
	case IdentifierMultihashSKEIN1024248:
		return uint64(HeaderMultihashSKEIN1024248)
	case IdentifierMultihashSKEIN1024256:
		return uint64(HeaderMultihashSKEIN1024256)
	case IdentifierMultihashSKEIN1024264:
		return uint64(HeaderMultihashSKEIN1024264)
	case IdentifierMultihashSKEIN1024272:
		return uint64(HeaderMultihashSKEIN1024272)
	case IdentifierMultihashSKEIN1024280:
		return uint64(HeaderMultihashSKEIN1024280)
	case IdentifierMultihashSKEIN1024288:
		return uint64(HeaderMultihashSKEIN1024288)
	case IdentifierMultihashSKEIN1024296:
		return uint64(HeaderMultihashSKEIN1024296)
	case IdentifierMultihashSKEIN1024304:
		return uint64(HeaderMultihashSKEIN1024304)
	case IdentifierMultihashSKEIN1024312:
		return uint64(HeaderMultihashSKEIN1024312)
	case IdentifierMultihashSKEIN1024320:
		return uint64(HeaderMultihashSKEIN1024320)
	case IdentifierMultihashSKEIN1024328:
		return uint64(HeaderMultihashSKEIN1024328)
	case IdentifierMultihashSKEIN1024336:
		return uint64(HeaderMultihashSKEIN1024336)
	case IdentifierMultihashSKEIN1024344:
		return uint64(HeaderMultihashSKEIN1024344)
	case IdentifierMultihashSKEIN1024352:
		return uint64(HeaderMultihashSKEIN1024352)
	case IdentifierMultihashSKEIN1024360:
		return uint64(HeaderMultihashSKEIN1024360)
	case IdentifierMultihashSKEIN1024368:
		return uint64(HeaderMultihashSKEIN1024368)
	case IdentifierMultihashSKEIN1024376:
		return uint64(HeaderMultihashSKEIN1024376)
	case IdentifierMultihashSKEIN1024384:
		return uint64(HeaderMultihashSKEIN1024384)
	case IdentifierMultihashSKEIN1024392:
		return uint64(HeaderMultihashSKEIN1024392)
	case IdentifierMultihashSKEIN1024400:
		return uint64(HeaderMultihashSKEIN1024400)
	case IdentifierMultihashSKEIN1024408:
		return uint64(HeaderMultihashSKEIN1024408)
	case IdentifierMultihashSKEIN1024416:
		return uint64(HeaderMultihashSKEIN1024416)
	case IdentifierMultihashSKEIN1024424:
		return uint64(HeaderMultihashSKEIN1024424)
	case IdentifierMultihashSKEIN1024432:
		return uint64(HeaderMultihashSKEIN1024432)
	case IdentifierMultihashSKEIN1024440:
		return uint64(HeaderMultihashSKEIN1024440)
	case IdentifierMultihashSKEIN1024448:
		return uint64(HeaderMultihashSKEIN1024448)
	case IdentifierMultihashSKEIN1024456:
		return uint64(HeaderMultihashSKEIN1024456)
	case IdentifierMultihashSKEIN1024464:
		return uint64(HeaderMultihashSKEIN1024464)
	case IdentifierMultihashSKEIN1024472:
		return uint64(HeaderMultihashSKEIN1024472)
	case IdentifierMultihashSKEIN1024480:
		return uint64(HeaderMultihashSKEIN1024480)
	case IdentifierMultihashSKEIN1024488:
		return uint64(HeaderMultihashSKEIN1024488)
	case IdentifierMultihashSKEIN1024496:
		return uint64(HeaderMultihashSKEIN1024496)
	case IdentifierMultihashSKEIN1024504:
		return uint64(HeaderMultihashSKEIN1024504)
	case IdentifierMultihashSKEIN1024512:
		return uint64(HeaderMultihashSKEIN1024512)
	case IdentifierMultihashSKEIN1024520:
		return uint64(HeaderMultihashSKEIN1024520)
	case IdentifierMultihashSKEIN1024528:
		return uint64(HeaderMultihashSKEIN1024528)
	case IdentifierMultihashSKEIN1024536:
		return uint64(HeaderMultihashSKEIN1024536)
	case IdentifierMultihashSKEIN1024544:
		return uint64(HeaderMultihashSKEIN1024544)
	case IdentifierMultihashSKEIN1024552:
		return uint64(HeaderMultihashSKEIN1024552)
	case IdentifierMultihashSKEIN1024560:
		return uint64(HeaderMultihashSKEIN1024560)
	case IdentifierMultihashSKEIN1024568:
		return uint64(HeaderMultihashSKEIN1024568)
	case IdentifierMultihashSKEIN1024576:
		return uint64(HeaderMultihashSKEIN1024576)
	case IdentifierMultihashSKEIN1024584:
		return uint64(HeaderMultihashSKEIN1024584)
	case IdentifierMultihashSKEIN1024592:
		return uint64(HeaderMultihashSKEIN1024592)
	case IdentifierMultihashSKEIN1024600:
		return uint64(HeaderMultihashSKEIN1024600)
	case IdentifierMultihashSKEIN1024608:
		return uint64(HeaderMultihashSKEIN1024608)
	case IdentifierMultihashSKEIN1024616:
		return uint64(HeaderMultihashSKEIN1024616)
	case IdentifierMultihashSKEIN1024624:
		return uint64(HeaderMultihashSKEIN1024624)
	case IdentifierMultihashSKEIN1024632:
		return uint64(HeaderMultihashSKEIN1024632)
	case IdentifierMultihashSKEIN1024640:
		return uint64(HeaderMultihashSKEIN1024640)
	case IdentifierMultihashSKEIN1024648:
		return uint64(HeaderMultihashSKEIN1024648)
	case IdentifierMultihashSKEIN1024656:
		return uint64(HeaderMultihashSKEIN1024656)
	case IdentifierMultihashSKEIN1024664:
		return uint64(HeaderMultihashSKEIN1024664)
	case IdentifierMultihashSKEIN1024672:
		return uint64(HeaderMultihashSKEIN1024672)
	case IdentifierMultihashSKEIN1024680:
		return uint64(HeaderMultihashSKEIN1024680)
	case IdentifierMultihashSKEIN1024688:
		return uint64(HeaderMultihashSKEIN1024688)
	case IdentifierMultihashSKEIN1024696:
		return uint64(HeaderMultihashSKEIN1024696)
	case IdentifierMultihashSKEIN1024704:
		return uint64(HeaderMultihashSKEIN1024704)
	case IdentifierMultihashSKEIN1024712:
		return uint64(HeaderMultihashSKEIN1024712)
	case IdentifierMultihashSKEIN1024720:
		return uint64(HeaderMultihashSKEIN1024720)
	case IdentifierMultihashSKEIN1024728:
		return uint64(HeaderMultihashSKEIN1024728)
	case IdentifierMultihashSKEIN1024736:
		return uint64(HeaderMultihashSKEIN1024736)
	case IdentifierMultihashSKEIN1024744:
		return uint64(HeaderMultihashSKEIN1024744)
	case IdentifierMultihashSKEIN1024752:
		return uint64(HeaderMultihashSKEIN1024752)
	case IdentifierMultihashSKEIN1024760:
		return uint64(HeaderMultihashSKEIN1024760)
	case IdentifierMultihashSKEIN1024768:
		return uint64(HeaderMultihashSKEIN1024768)
	case IdentifierMultihashSKEIN1024776:
		return uint64(HeaderMultihashSKEIN1024776)
	case IdentifierMultihashSKEIN1024784:
		return uint64(HeaderMultihashSKEIN1024784)
	case IdentifierMultihashSKEIN1024792:
		return uint64(HeaderMultihashSKEIN1024792)
	case IdentifierMultihashSKEIN1024800:
		return uint64(HeaderMultihashSKEIN1024800)
	case IdentifierMultihashSKEIN1024808:
		return uint64(HeaderMultihashSKEIN1024808)
	case IdentifierMultihashSKEIN1024816:
		return uint64(HeaderMultihashSKEIN1024816)
	case IdentifierMultihashSKEIN1024824:
		return uint64(HeaderMultihashSKEIN1024824)
	case IdentifierMultihashSKEIN1024832:
		return uint64(HeaderMultihashSKEIN1024832)
	case IdentifierMultihashSKEIN1024840:
		return uint64(HeaderMultihashSKEIN1024840)
	case IdentifierMultihashSKEIN1024848:
		return uint64(HeaderMultihashSKEIN1024848)
	case IdentifierMultihashSKEIN1024856:
		return uint64(HeaderMultihashSKEIN1024856)
	case IdentifierMultihashSKEIN1024864:
		return uint64(HeaderMultihashSKEIN1024864)
	case IdentifierMultihashSKEIN1024872:
		return uint64(HeaderMultihashSKEIN1024872)
	case IdentifierMultihashSKEIN1024880:
		return uint64(HeaderMultihashSKEIN1024880)
	case IdentifierMultihashSKEIN1024888:
		return uint64(HeaderMultihashSKEIN1024888)
	case IdentifierMultihashSKEIN1024896:
		return uint64(HeaderMultihashSKEIN1024896)
	case IdentifierMultihashSKEIN1024904:
		return uint64(HeaderMultihashSKEIN1024904)
	case IdentifierMultihashSKEIN1024912:
		return uint64(HeaderMultihashSKEIN1024912)
	case IdentifierMultihashSKEIN1024920:
		return uint64(HeaderMultihashSKEIN1024920)
	case IdentifierMultihashSKEIN1024928:
		return uint64(HeaderMultihashSKEIN1024928)
	case IdentifierMultihashSKEIN1024936:
		return uint64(HeaderMultihashSKEIN1024936)
	case IdentifierMultihashSKEIN1024944:
		return uint64(HeaderMultihashSKEIN1024944)
	case IdentifierMultihashSKEIN1024952:
		return uint64(HeaderMultihashSKEIN1024952)
	case IdentifierMultihashSKEIN1024960:
		return uint64(HeaderMultihashSKEIN1024960)
	case IdentifierMultihashSKEIN1024968:
		return uint64(HeaderMultihashSKEIN1024968)
	case IdentifierMultihashSKEIN1024976:
		return uint64(HeaderMultihashSKEIN1024976)
	case IdentifierMultihashSKEIN1024984:
		return uint64(HeaderMultihashSKEIN1024984)
	case IdentifierMultihashSKEIN1024992:
		return uint64(HeaderMultihashSKEIN1024992)
	case IdentifierMultihashSKEIN10241000:
		return uint64(HeaderMultihashSKEIN10241000)
	case IdentifierMultihashSKEIN10241008:
		return uint64(HeaderMultihashSKEIN10241008)
	case IdentifierMultihashSKEIN10241016:
		return uint64(HeaderMultihashSKEIN10241016)
	case IdentifierMultihashSKEIN10241024:
		return uint64(HeaderMultihashSKEIN10241024)
	case IdentifierHashXXH32:
		return uint64(HeaderHashXXH32)
	case IdentifierHashXXH64:
		return uint64(HeaderHashXXH64)
	case IdentifierHashXXH364:
		return uint64(HeaderHashXXH364)
	case IdentifierHashXXH3128:
		return uint64(HeaderHashXXH3128)
	case IdentifierMultihashPOSEIDONBLS12381A2FC1:
		return uint64(HeaderMultihashPOSEIDONBLS12381A2FC1)
	case IdentifierMultihashPOSEIDONBLS12381A2FC1SC:
		return uint64(HeaderMultihashPOSEIDONBLS12381A2FC1SC)
	case IdentifierIpldRDFC1:
		return uint64(HeaderIpldRDFC1)
	case IdentifierSerializationSSZ:
		return uint64(HeaderSerializationSSZ)
	case IdentifierMultihashSSZSHA2256BMT:
		return uint64(HeaderMultihashSSZSHA2256BMT)
	case IdentifierIpldJSONJCS:
		return uint64(HeaderIpldJSONJCS)
	case IdentifierSofthashISCC:
		return uint64(HeaderSofthashISCC)
	case IdentifierZeroxcertZEROXCERTIMPRINT256:
		return uint64(HeaderZeroxcertZEROXCERTIMPRINT256)
	case IdentifierVarsigNONSTANDARDSIG:
		return uint64(HeaderVarsigNONSTANDARDSIG)
	case IdentifierVarsigES256K:
		return uint64(HeaderVarsigES256K)
	case IdentifierVarsigBLS12381G1SIG:
		return uint64(HeaderVarsigBLS12381G1SIG)
	case IdentifierVarsigBLS12381G2SIG:
		return uint64(HeaderVarsigBLS12381G2SIG)
	case IdentifierVarsigEDDSA:
		return uint64(HeaderVarsigEDDSA)
	case IdentifierVarsigEIP191:
		return uint64(HeaderVarsigEIP191)
	case IdentifierKeyJWKJCSPUB:
		return uint64(HeaderKeyJWKJCSPUB)
	case IdentifierFilecoinFILCOMMITMENTUNSEALED:
		return uint64(HeaderFilecoinFILCOMMITMENTUNSEALED)
	case IdentifierFilecoinFILCOMMITMENTSEALED:
		return uint64(HeaderFilecoinFILCOMMITMENTSEALED)
	case IdentifierMultiaddrPLAINTEXTV2:
		return uint64(HeaderMultiaddrPLAINTEXTV2)
	case IdentifierHolochainHOLOCHAINADRV0:
		return uint64(HeaderHolochainHOLOCHAINADRV0)
	case IdentifierHolochainHOLOCHAINADRV1:
		return uint64(HeaderHolochainHOLOCHAINADRV1)
	case IdentifierHolochainHOLOCHAINKEYV0:
		return uint64(HeaderHolochainHOLOCHAINKEYV0)
	case IdentifierHolochainHOLOCHAINKEYV1:
		return uint64(HeaderHolochainHOLOCHAINKEYV1)
	case IdentifierHolochainHOLOCHAINSIGV0:
		return uint64(HeaderHolochainHOLOCHAINSIGV0)
	case IdentifierHolochainHOLOCHAINSIGV1:
		return uint64(HeaderHolochainHOLOCHAINSIGV1)
	case IdentifierNamespaceSKYNETNS:
		return uint64(HeaderNamespaceSKYNETNS)
	case IdentifierNamespaceARWEAVENS:
		return uint64(HeaderNamespaceARWEAVENS)
	case IdentifierNamespaceSUBSPACENS:
		return uint64(HeaderNamespaceSUBSPACENS)
	case IdentifierNamespaceKUMANDRANS:
		return uint64(HeaderNamespaceKUMANDRANS)
	case IdentifierVarsigES256:
		return uint64(HeaderVarsigES256)
	case IdentifierVarsigES284:
		return uint64(HeaderVarsigES284)
	case IdentifierVarsigES512:
		return uint64(HeaderVarsigES512)
	case IdentifierVarsigRS256:
		return uint64(HeaderVarsigRS256)
	case IdentifierMultiaddrSCION:
		return uint64(HeaderMultiaddrSCION)
	default:
		return uint64(HeaderUNKNOWN)
	}
}
