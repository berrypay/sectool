################################################################################
# Project: BerryPay Application Security Tool Set                              #
# Filename: /Makefile                                                          #
# Created Date: Friday September 1st 2023 11:17:29 +0800                       #
# Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)                   #
# Company: BerryPay (M) Sdn. Bhd.                                              #
# --------------------------------------                                       #
# Last Modified: Friday September 1st 2023 11:21:46 +0800                      #
# Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)              #
# --------------------------------------                                       #
# Copyright (c) 2023 BerryPay (M) Sdn. Bhd.                                    #
################################################################################

default: build

build:
	go build -o sectool main.go

clean:
	rm -f sectool

