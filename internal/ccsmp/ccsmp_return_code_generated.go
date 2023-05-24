// pubsubplus-go-client
//
// Copyright 2021-2023 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ccsmp

// Code generated by ccsmp_return_code_generator.go via go generate. DO NOT EDIT.

/*
#include "solclient/solClient.h"
*/
import "C"

const (
	// SolClientReturnCodeOk: The API call was successful.
	SolClientReturnCodeOk SolClientReturnCode = C.SOLCLIENT_OK
	// SolClientReturnCodeWouldBlock: The API call would block, but non-blocking was requested.
	SolClientReturnCodeWouldBlock SolClientReturnCode = C.SOLCLIENT_WOULD_BLOCK
	// SolClientReturnCodeInProgress: An API call is in progress (non-blocking mode).
	SolClientReturnCodeInProgress SolClientReturnCode = C.SOLCLIENT_IN_PROGRESS
	// SolClientReturnCodeNotReady: The API could not complete as an object is not ready (for example, the Session is not connected).
	SolClientReturnCodeNotReady SolClientReturnCode = C.SOLCLIENT_NOT_READY
	// SolClientReturnCodeEos: A getNext on a structured container returned End-of-Stream.
	SolClientReturnCodeEos SolClientReturnCode = C.SOLCLIENT_EOS
	// SolClientReturnCodeNotFound: A get for a named field in a MAP was not found in the MAP.
	SolClientReturnCodeNotFound SolClientReturnCode = C.SOLCLIENT_NOT_FOUND
	// SolClientReturnCodeNoevent: solClient_context_processEventsWait returns this if wait is zero and there is no event to process
	SolClientReturnCodeNoevent SolClientReturnCode = C.SOLCLIENT_NOEVENT
	// SolClientReturnCodeIncomplete: The API call completed some, but not all, of the requested function.
	SolClientReturnCodeIncomplete SolClientReturnCode = C.SOLCLIENT_INCOMPLETE
	// SolClientReturnCodeRollback: solClient_transactedSession_commit returns this when the transaction has been rolled back.
	SolClientReturnCodeRollback SolClientReturnCode = C.SOLCLIENT_ROLLBACK
	// SolClientReturnCodeFail: The API call failed.
	SolClientReturnCodeFail SolClientReturnCode = C.SOLCLIENT_FAIL
)
