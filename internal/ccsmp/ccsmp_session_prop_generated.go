// pubsubplus-go-client
//
// Copyright 2021-2022 Solace Corporation. All rights reserved.
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

// Code generated by ccsmp_session_prop_generator.go via go generate. DO NOT EDIT.

/*
#include "solclient/solClient.h"
*/
import "C"

const (
	// SolClientSessionPropAckEventModePerMsg: In this mode,  a session event ::SOLCLIENT_SESSION_EVENT_ACKNOWLEDGEMENT acknowledges a single message.
	SolClientSessionPropAckEventModePerMsg = C.SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE_PER_MSG
	// SolClientSessionPropAckEventModeWindowed: In this mode, a session event ::SOLCLIENT_SESSION_EVENT_ACKNOWLEDGEMENT is a ranged acknowledgment. It acknowledges the message received as well as all preceding messages.
	SolClientSessionPropAckEventModeWindowed = C.SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE_WINDOWED
	// SolClientSessionPropUsername: The username required for authentication. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_USERNAME
	SolClientSessionPropUsername = C.SOLCLIENT_SESSION_PROP_USERNAME
	// SolClientSessionPropPassword: The password required for authentication. May be set as an environment variable (See @ref SessionProps). Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_PASSWORD
	SolClientSessionPropPassword = C.SOLCLIENT_SESSION_PROP_PASSWORD
	// SolClientSessionPropHost: The IPv4 or IPv6 address or host name to connect to. @ref host-list "Multiple entries" (up to ::SOLCLIENT_SESSION_PROP_MAX_HOSTS) are allowed, separated by commas. @ref host-entry "The entry for the SOLCLIENT_SESSION_PROP_HOST property should provide a protocol, host, and port". See @ref host-list "Configuring Multiple Hosts for Redundancy and Failover" for a discussion of Guaranteed Messaging considerations. May be set as an environment variable (See @ref SessionProps). Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_HOST
	SolClientSessionPropHost = C.SOLCLIENT_SESSION_PROP_HOST
	// SolClientSessionPropPort: Deprecated. While still supported, the port number can also now be specified as part of the host in ::SOLCLIENT_SESSION_PROP_HOST (for example, "hostname:55555"). In general, port numbers are not needed except in special situations, because the API chooses the correct port to connect to the broker. If ::SOLCLIENT_SESSION_PROP_PORT is set, this port number is used for all entries in ::SOLCLIENT_SESSION_PROP_HOST that do not explicitly specify port. The port number to connect to. The valid range is 1..65535. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_PORT or ::SOLCLIENT_SESSION_PROP_DEFAULT_PORT_COMPRESSION, based on setting of ::SOLCLIENT_SESSION_PROP_COMPRESSION_LEVEL
	SolClientSessionPropPort = C.SOLCLIENT_SESSION_PROP_PORT
	// SolClientSessionPropBufferSize: The maximum amount of messages to buffer (in bytes) when the TCP session is flow controlled (see \ref message-buffer "Message Buffer Size Configuration"). The valid range is > 0. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_BUFFER_SIZE
	SolClientSessionPropBufferSize = C.SOLCLIENT_SESSION_PROP_BUFFER_SIZE
	// SolClientSessionPropConnectBlocking: Use ::SOLCLIENT_PROP_ENABLE_VAL to enable blocking connect operation. A blocking connect operation suspends until the Session is successfully connected, including restoring all remembered subscriptions if ::SOLCLIENT_SESSION_PROP_REAPPLY_SUBSCRIPTIONS is enabled. Otherwise solClient_session_connect() returns SOLCLIENT_IN_PROGRESS. @see @ref blocking-context "Threading Effects on Blocking Modes" for a discussion of blocking operation in the Context thread. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_CONNECT_BLOCKING
	SolClientSessionPropConnectBlocking = C.SOLCLIENT_SESSION_PROP_CONNECT_BLOCKING
	// SolClientSessionPropSendBlocking: Use ::SOLCLIENT_PROP_ENABLE_VAL to enable blocking send operation. A blocking send operation suspends when the Session is transport flow controlled, otherwise the send operation returns SOLCLIENT_WOULD_BLOCK. Successful return from a blocking send operation only means the message has been accepted by the transport, it does not guarantee the message has been processed by the broker. For the latter you must used Guaranteed Message Delivery mode and wait for the session event (::SOLCLIENT_SESSION_EVENT_ACKNOWLEDGEMENT)  that acknowledges the message. @see @ref blocking-context "Threading Effects on Blocking Modes" for a discussion of blocking operation in the Context thread. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SEND_BLOCKING
	SolClientSessionPropSendBlocking = C.SOLCLIENT_SESSION_PROP_SEND_BLOCKING
	// SolClientSessionPropSubscribeBlocking: Use ::SOLCLIENT_PROP_ENABLE_VAL to enable blocking subscribe/unsubscribe operation. A blocking subscribe operation will suspend when the Session is transport flow controlled, otherwise the subscribe operation returns SOLCLIENT_WOULD_BLOCK. A successful return from a blocking subscribe operation only means the subscription has been accepted by the transport, it does not guarantee the subscription has been processed by the broker. For the latter you must use a confirmed operation (see ::SOLCLIENT_SUBSCRIBE_FLAGS_WAITFORCONFIRM)  @see @ref blocking-context "Threading Effects on Blocking Modes" for a discussion of blocking operation in the Context thread. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SUBSCRIBE_BLOCKING
	SolClientSessionPropSubscribeBlocking = C.SOLCLIENT_SESSION_PROP_SUBSCRIBE_BLOCKING
	// SolClientSessionPropBlockWhileConnecting: Use ::SOLCLIENT_PROP_ENABLE_VAL to block the calling thread on operations such as sending a message, subscribing, or unsubscribing when the Session is being connected or reconnected. The operation must already be blocking (see ::SOLCLIENT_SESSION_PROP_SEND_BLOCKING and ::SOLCLIENT_SESSION_PROP_SUBSCRIBE_BLOCKING). Otherwise, ::SOLCLIENT_NOT_READY is returned if the Session is being connected. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_BLOCK_WHILE_CONNECTING
	SolClientSessionPropBlockWhileConnecting = C.SOLCLIENT_SESSION_PROP_BLOCK_WHILE_CONNECTING
	// SolClientSessionPropBlockingWriteTimeoutMs: The timeout period (in milliseconds) for blocking write operation. The valid range is > 0. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_BLOCKING_WRITE_TIMEOUT_MS
	SolClientSessionPropBlockingWriteTimeoutMs = C.SOLCLIENT_SESSION_PROP_BLOCKING_WRITE_TIMEOUT_MS
	// SolClientSessionPropConnectTimeoutMs: The timeout period (in milliseconds) for a connect operation to a given host (per host). The valid range is > 0. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_CONNECT_TIMEOUT_MS
	SolClientSessionPropConnectTimeoutMs = C.SOLCLIENT_SESSION_PROP_CONNECT_TIMEOUT_MS
	// SolClientSessionPropSubconfirmTimeoutMs: The timeout period (in milliseconds) for subscription confirm (add or remove). The valid range is >= 1000. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SUBCONFIRM_TIMEOUT_MS
	SolClientSessionPropSubconfirmTimeoutMs = C.SOLCLIENT_SESSION_PROP_SUBCONFIRM_TIMEOUT_MS
	// SolClientSessionPropIgnoreDupSubscriptionError: Use ::SOLCLIENT_PROP_ENABLE_VAL to ignore errors for duplicate subscription/topic on subscribe or subscription not found errors on unsubscribe. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_IGNORE_DUP_SUBSCRIPTION_ERROR
	SolClientSessionPropIgnoreDupSubscriptionError = C.SOLCLIENT_SESSION_PROP_IGNORE_DUP_SUBSCRIPTION_ERROR
	// SolClientSessionPropTCPNodelay: Use ::SOLCLIENT_PROP_ENABLE_VAL to enable TCP no delay. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_TCP_NODELAY
	SolClientSessionPropTCPNodelay = C.SOLCLIENT_SESSION_PROP_TCP_NODELAY
	// SolClientSessionPropSocketSendBufSize: The value for the socket send buffer size (in bytes). 0 indicates do not set and leave at operating system default. The valid range is 0 or >= 1024. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SOCKET_SEND_BUF_SIZE.
	SolClientSessionPropSocketSendBufSize = C.SOLCLIENT_SESSION_PROP_SOCKET_SEND_BUF_SIZE
	// SolClientSessionPropSocketRcvBufSize: The value for socket receive buffer size  (in bytes). 0 indicates do not set and leave at operating system default. The valid range is 0 or >= 1024. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SOCKET_RCV_BUF_SIZE
	SolClientSessionPropSocketRcvBufSize = C.SOLCLIENT_SESSION_PROP_SOCKET_RCV_BUF_SIZE
	// SolClientSessionPropKeepAliveIntMs: The amount of time (in milliseconds) to wait between sending out Keep-Alive messages. Typically, this feature should be enabled for message receivers. Use 0 to disable Keep-Alives (0 is required before broker release 4.2). The valid range is 0 (disabled) or >= 50. Default:  ::SOLCLIENT_SESSION_PROP_DEFAULT_KEEP_ALIVE_INT_MS
	SolClientSessionPropKeepAliveIntMs = C.SOLCLIENT_SESSION_PROP_KEEP_ALIVE_INT_MS
	// SolClientSessionPropKeepAliveLimit: The maximum number of consecutive Keep-Alive messages that can be sent without receiving a response before the connection is closed by the API. The valid range is >= 3. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_KEEP_ALIVE_LIMIT
	SolClientSessionPropKeepAliveLimit = C.SOLCLIENT_SESSION_PROP_KEEP_ALIVE_LIMIT
	// SolClientSessionPropApplicationDescription: A string that uniquely describes the application instance. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_APPLICATION_DESCRIPTION
	SolClientSessionPropApplicationDescription = C.SOLCLIENT_SESSION_PROP_APPLICATION_DESCRIPTION
	// SolClientSessionPropClientMode: Deprecated. The CCSMP API detects the broker capabilities, so it is no longer necessary to specify to use 'clientMode' or not. This property is ignored when specified.
	SolClientSessionPropClientMode = C.SOLCLIENT_SESSION_PROP_CLIENT_MODE
	// SolClientSessionPropBindIP: (Optional) The hostname or IP address of the machine on which the application is running. On a multihomed machine, it is strongly recommended to provide this parameter to ensure that the API uses the correct network interface at Session connect time. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_BIND_IP
	SolClientSessionPropBindIP = C.SOLCLIENT_SESSION_PROP_BIND_IP
	// SolClientSessionPropPubWindowSize: The publisher window size for Guaranteed messages. The Guaranteed Message Publish Window Size property limits the maximum number of messages that can be published before the API must receive an acknowledgment from the broker. The valid range is 1..255, or 0 to disable publishing Guaranteed messages. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_PUB_WINDOW_SIZE
	SolClientSessionPropPubWindowSize = C.SOLCLIENT_SESSION_PROP_PUB_WINDOW_SIZE
	// SolClientSessionPropPubAckTimer: The duration of publisher acknowledgment timer (in milliseconds). When a published message is not acknowledged within the time specified for this timer, the API automatically retransmits the message. There is no limit on the number of retransmissions for any message. However, while the API is resending, applications can become flow controlled. The flow control behavior is controlled by ::SOLCLIENT_SESSION_PROP_SEND_BLOCKING and ::SOLCLIENT_SESSION_PROP_BLOCKING_WRITE_TIMEOUT_MS. The valid range is 20..60000. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_PUB_ACK_TIMER
	SolClientSessionPropPubAckTimer = C.SOLCLIENT_SESSION_PROP_PUB_ACK_TIMER
	// SolClientSessionPropVpnName: The name of the Message VPN to attempt to join when connecting to an broker running SolOS-TR. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_VPN_NAME
	SolClientSessionPropVpnName = C.SOLCLIENT_SESSION_PROP_VPN_NAME
	// SolClientSessionPropVpnNameInUse: A read-only Session property that indicates which Message VPN the Session is connected to. When not connected, an empty string is returned.
	SolClientSessionPropVpnNameInUse = C.SOLCLIENT_SESSION_PROP_VPN_NAME_IN_USE
	// SolClientSessionPropClientName: The Session client name that is used during client login to create a unique Session. An empty string causes a unique client name to be generated automatically. If specified, it must be a valid Topic name, and a maximum of 160 bytes in length. For all brokers (SolOS-TR or SolOS-CR) the SOLCLIENT_SESSION_PROP_CLIENT_NAME is also used to uniquely identify the sender in a message's senderId field if ::SOLCLIENT_SESSION_PROP_GENERATE_SENDER_ID is set. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_CLIENT_NAME
	SolClientSessionPropClientName = C.SOLCLIENT_SESSION_PROP_CLIENT_NAME
	// SolClientSessionPropCompressionLevel: Enables messages to be compressed with ZLIB before transmission and decompressed on receive. The valid range is 0 (off) or 1..9, where 1 is less compression (fastest) and 9 is most compression (slowest). Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_COMPRESSION_LEVEL
	SolClientSessionPropCompressionLevel = C.SOLCLIENT_SESSION_PROP_COMPRESSION_LEVEL
	// SolClientSessionPropGenerateRcvTimestamps: When enabled, a receive timestamp is recorded for each message and passed to the application callback in the rxCallbackInfo_t structure. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_GENERATE_RCV_TIMESTAMPS
	SolClientSessionPropGenerateRcvTimestamps = C.SOLCLIENT_SESSION_PROP_GENERATE_RCV_TIMESTAMPS
	// SolClientSessionPropGenerateSendTimestamps: When enabled, a send timestamp is automatically included (if not already present) in the Solace-defined fields for each message sent. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_GENERATE_SEND_TIMESTAMPS
	SolClientSessionPropGenerateSendTimestamps = C.SOLCLIENT_SESSION_PROP_GENERATE_SEND_TIMESTAMPS
	// SolClientSessionPropGenerateSenderID: When enabled, a sender ID is automatically included (if not already present) in the Solace-defined fields for each message sent. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_GENERATE_SENDER_ID
	SolClientSessionPropGenerateSenderID = C.SOLCLIENT_SESSION_PROP_GENERATE_SENDER_ID
	// SolClientSessionPropGenerateSequenceNumber: When enabled, a sequence number is automatically included (if not already present) in the Solace-defined fields for each message sent. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_GENERATE_SEQUENCE_NUMBER
	SolClientSessionPropGenerateSequenceNumber = C.SOLCLIENT_SESSION_PROP_GENERATE_SEQUENCE_NUMBER
	// SolClientSessionPropConnectRetriesPerHost:  When using a host list, this property defines how many times to try to connect or reconnect to a single host before moving to the next host in the list. A value of 0 (the default) means make a single connection attempt (that is, 0 retries). A value of -1 means attempt an infinite number of reconnect retries (that is, the API will only try to connect or reconnect to first host listed.) NOTE: This property works in conjunction with the connect and reconnect retries Session properties; it does not replace them.
	SolClientSessionPropConnectRetriesPerHost = C.SOLCLIENT_SESSION_PROP_CONNECT_RETRIES_PER_HOST
	// SolClientSessionPropConnectRetries: How many times to try to connect to the host broker (or list of broker) during connection setup. Zero means no automatic connection retries (that is, try once and give up). -1 means try to connect forever. The default valid range is >= -1.
	SolClientSessionPropConnectRetries = C.SOLCLIENT_SESSION_PROP_CONNECT_RETRIES
	// SolClientSessionPropReconnectRetries: How many times to retry to reconnect to the host broker (or list of broker) after a connected Session goes down. Zero means no automatic reconnection attempts. -1 means try to reconnect forever. The default valid range is >= -1.
	SolClientSessionPropReconnectRetries = C.SOLCLIENT_SESSION_PROP_RECONNECT_RETRIES
	// SolClientSessionPropReconnectRetryWaitMs: How much time (in ms) to wait between each attempt to connect or reconnect to a host. If a connect or reconnect attempt to host is not successful, the API waits for the amount of time set for SOLCLIENT_SESSION_PROP_RECONNECT_RETRY_WAIT_MS, and then makes another connect or reconnect attempt. SESSION_CONNECT_RETRIES_PER_HOST sets how many connection or reconnection attempts can be made before moving
	SolClientSessionPropReconnectRetryWaitMs = C.SOLCLIENT_SESSION_PROP_RECONNECT_RETRY_WAIT_MS
	// SolClientSessionPropUserID: A read-only informational string providing information about the application, such as the name of operating system user that is running the application, the hostname, and the PID of the application.
	SolClientSessionPropUserID = C.SOLCLIENT_SESSION_PROP_USER_ID
	// SolClientSessionPropP2pinboxInUse: A read-only informational string that indicates the default reply-to destination that is used when a request message is sent that does not have a reply-to destination specified. See solClient_session_sendRequest(). This parameter is only valid when the Session is connected.
	SolClientSessionPropP2pinboxInUse = C.SOLCLIENT_SESSION_PROP_P2PINBOX_IN_USE
	// SolClientSessionPropReapplySubscriptions: Use ::SOLCLIENT_PROP_ENABLE_VAL to have the API remember subscriptions and reapply them upon a Session reconnect. Reapply subscriptions will only apply direct topic subscriptions upon a Session reconnect. It will not reapply topic subscriptions on durable and non-durable endpoints. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_REAPPLY_SUBSCRIPTIONS
	SolClientSessionPropReapplySubscriptions = C.SOLCLIENT_SESSION_PROP_REAPPLY_SUBSCRIPTIONS
	// SolClientSessionPropTopicDispatch: Use ::SOLCLIENT_PROP_ENABLE_VAL to have the API dispatch messages based on Topic (see @ref topic-dispatch). Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_TOPIC_DISPATCH
	SolClientSessionPropTopicDispatch = C.SOLCLIENT_SESSION_PROP_TOPIC_DISPATCH
	// SolClientSessionPropProvisionTimeoutMs: Maximum amount of time (in milliseconds) to wait for a provision command (create or delete an endpoint)
	SolClientSessionPropProvisionTimeoutMs = C.SOLCLIENT_SESSION_PROP_PROVISION_TIMEOUT_MS
	// SolClientSessionPropCalculateMessageExpiration: If this property is true and time-to-live (::solClient_msg_setTimeToLive()) has a positive value in a message, the expiration time is calculated when the message is sent or received and can be retrieved with ::solClient_msg_getExpiration.
	SolClientSessionPropCalculateMessageExpiration = C.SOLCLIENT_SESSION_PROP_CALCULATE_MESSAGE_EXPIRATION
	// SolClientSessionPropVirtualRouterName: A read-only property that indicates the connected broker's virtual router name. Appliance endpoint and destination names created with a virtual router name are valid for use with that broker, or to address destinations on remote brokers (in a multiple-broker network) when publishing messages. Applications requiring the virtual router name do not need to poll this property every time it is required, and they may cache the name. Applications should query the name once after connecting the Session, and again after a reconnect operation reports the ::SOLCLIENT_SESSION_EVENT_VIRTUAL_ROUTER_NAME_CHANGED event. Prior to connecting, an empty string is returned.
	SolClientSessionPropVirtualRouterName = C.SOLCLIENT_SESSION_PROP_VIRTUAL_ROUTER_NAME
	// SolClientSessionPropNoLocal: If this property is true, messages published on the Session cannot be received on the same Session even if the client has a subscription that matches the published topic. If this restriction is requested, and the broker does not have No Local support, the Session connect will fail with subcode ::SOLCLIENT_SUBCODE_NO_LOCAL_NOT_SUPPORTED.
	SolClientSessionPropNoLocal = C.SOLCLIENT_SESSION_PROP_NO_LOCAL
	// SolClientSessionPropAdPubRouterWindowedAck: When disabled, initiate a window size of 1 to broker, but do not wait for acknowledgments before transmitting up to the actual window size. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_AD_PUB_ROUTER_WINDOWED_ACK
	SolClientSessionPropAdPubRouterWindowedAck = C.SOLCLIENT_SESSION_PROP_AD_PUB_ROUTER_WINDOWED_ACK
	// SolClientSessionPropModifypropTimeoutMs: Maximum amount of time (in milliseconds) to wait for session property modification. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_MODIFYPROP_TIMEOUT_MS
	SolClientSessionPropModifypropTimeoutMs = C.SOLCLIENT_SESSION_PROP_MODIFYPROP_TIMEOUT_MS
	// SolClientSessionPropAckEventMode: This property specifies if a session event ::SOLCLIENT_SESSION_EVENT_ACKNOWLEDGEMENT acknowledges a single message (see ::SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE_PER_MSG) or a range of messages (see ::SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE_WINDOWED).  Default: ::SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE_PER_MSG. \n Setting this property to ::SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE_WINDOWED will not affect RejectedMessageError events, they will still be emitted on a per message basis.
	SolClientSessionPropAckEventMode = C.SOLCLIENT_SESSION_PROP_ACK_EVENT_MODE
	// SolClientSessionPropSslExcludedProtocols: This property specifies a comma separated list of excluded SSL protocol(s). Valid SSL protocols are 'SSLv3', 'TLSv1', 'TLSv1.1', 'TLSv1.2'. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SSL_EXCLUDED_PROTOCOLS.
	SolClientSessionPropSslExcludedProtocols = C.SOLCLIENT_SESSION_PROP_SSL_EXCLUDED_PROTOCOLS
	// SolClientSessionPropSslValidateCertificate: This is used to specify whether the API should validate server certificates with certificates in the truststore. When disabled, validation of the certificate date, certificate host, and checking the common name list is also disabled. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SSL_VALIDATE_CERTIFICATE.<p> See also @ref certificate-validation
	SolClientSessionPropSslValidateCertificate = C.SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE
	// SolClientSessionPropOpensslSecurityLevel: A number from 0-5 passed to the openSsl library as the security level. Use 0 with caution, it is less secure than the default, which is 1.
	SolClientSessionPropOpensslSecurityLevel = C.SOLCLIENT_SESSION_PROP_OPENSSL_SECURITY_LEVEL
	// SolClientSessionPropSslClientCertificateFile: This property specifies the client certificate file name.
	SolClientSessionPropSslClientCertificateFile = C.SOLCLIENT_SESSION_PROP_SSL_CLIENT_CERTIFICATE_FILE
	// SolClientSessionPropSslClientPrivateKeyFile: This property specifies the client private key file name.
	SolClientSessionPropSslClientPrivateKeyFile = C.SOLCLIENT_SESSION_PROP_SSL_CLIENT_PRIVATE_KEY_FILE
	// SolClientSessionPropSslClientPrivateKeyFilePassword: This property specifies the password used to encrypt the client private key file.
	SolClientSessionPropSslClientPrivateKeyFilePassword = C.SOLCLIENT_SESSION_PROP_SSL_CLIENT_PRIVATE_KEY_FILE_PASSWORD
	// SolClientSessionPropSslConnectionDowngradeTo: This property specifies a transport protocol that SSL connection will be downgraded to after client authentication. Allowed transport protocol is "PLAIN_TEXT". May be combined with non-zero compression level to achieve compression without encryption.
	SolClientSessionPropSslConnectionDowngradeTo = C.SOLCLIENT_SESSION_PROP_SSL_CONNECTION_DOWNGRADE_TO
	// SolClientSessionPropInitialReceiveBufferSize: If not zero, the minimum starting size for the API receive buffer. Must be zero or >= 1024 and <=64*1024*1024
	SolClientSessionPropInitialReceiveBufferSize = C.SOLCLIENT_SESSION_PROP_INITIAL_RECEIVE_BUFFER_SIZE
	// SolClientSessionPropAuthenticationScheme: This property specifies the authentication scheme. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_AUTHENTICATION_SCHEME.
	SolClientSessionPropAuthenticationScheme = C.SOLCLIENT_SESSION_PROP_AUTHENTICATION_SCHEME
	// SolClientSessionPropKrbServiceName: This property specifies the first part of Kerberos Service Principal Name (SPN) of the form <i>ServiceName/Hostname\@REALM</i> (for Windows) or Host Based Service of the form <i>ServiceName\@Hostname</i> (for Linux and SunOS).
	SolClientSessionPropKrbServiceName = C.SOLCLIENT_SESSION_PROP_KRB_SERVICE_NAME
	// SolClientSessionPropUnbindFailAction: A property to define the behavior if an unbind-response is not received after an unbind-request (::solClient_flow_destroy()) is sent to the Solace Appliance. If this occurs it is possible that the endpoint may be still bound and unavailable until the session is terminated. In this occurrence the session can be configured to retry sending the unbind-request or to fail the session transport. If the sesion transport fails, the session will behave as defined by the ::SOLCLIENT_SESSION_PROP_RECONNECT_RETRIES configuration.  Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_UNBIND_FAIL_ACTION.
	SolClientSessionPropUnbindFailAction = C.SOLCLIENT_SESSION_PROP_UNBIND_FAIL_ACTION
	// SolClientSessionPropWebTransportProtocol: This property specifies a WEB Transport Protocol in the default WEB Transport Protocol downgrade list to use for the session connection.
	SolClientSessionPropWebTransportProtocol = C.SOLCLIENT_SESSION_PROP_WEB_TRANSPORT_PROTOCOL
	// SolClientSessionPropWebTransportProtocolInUse: Read-only property which returns the WEB Transport Protocol currently in use for web messaging. An empty string is returned when a session
	SolClientSessionPropWebTransportProtocolInUse = C.SOLCLIENT_SESSION_PROP_WEB_TRANSPORT_PROTOCOL_IN_USE
	// SolClientSessionPropWebTransportProtocolList: This property specifies a comma separated list of WEB Transport Protocols to use for session connection. The API will use the first one in
	SolClientSessionPropWebTransportProtocolList = C.SOLCLIENT_SESSION_PROP_WEB_TRANSPORT_PROTOCOL_LIST
	// SolClientSessionPropTransportProtocolDowngradeTimeoutMs: Specifies how long to wait (in milliseconds) for a login response before moving to the next available protocol in a user specified or the default WEB Transport Protocol downgradelist.
	SolClientSessionPropTransportProtocolDowngradeTimeoutMs = C.SOLCLIENT_SESSION_PROP_TRANSPORT_PROTOCOL_DOWNGRADE_TIMEOUT_MS
	// SolClientSessionPropGuaranteedWithWebTransport: Enables guaranteed messaging with web transport protocols. Refer to \ref transportProtocol for supported transport protocols. Only has effect when a session is established over a web messaging transport. Default ::SOLCLIENT_SESSION_PROP_DEFAULT_GUARANTEED_WITH_WEB_TRANSPORT
	SolClientSessionPropGuaranteedWithWebTransport = C.SOLCLIENT_SESSION_PROP_GUARANTEED_WITH_WEB_TRANSPORT
	// SolClientSessionPropGdReconnectFailAction: A property to define the behavior when the CCSMP API is unable
	SolClientSessionPropGdReconnectFailAction = C.SOLCLIENT_SESSION_PROP_GD_RECONNECT_FAIL_ACTION
	// SolClientSessionPropOauth2AccessToken: The OAUTH2 access token. When authentication scheme ::SOLCLIENT_SESSION_PROP_AUTHENTICATION_SCHEME_OAUTH2 is used at least one of ::SOLCLIENT_SESSION_PROP_OAUTH2_ACCESS_TOKEN and ::SOLCLIENT_SESSION_PROP_OIDC_ID_TOKEN must be set.
	SolClientSessionPropOauth2AccessToken = C.SOLCLIENT_SESSION_PROP_OAUTH2_ACCESS_TOKEN
	// SolClientSessionPropOauth2IssuerIdentifier: The optional Issuer identifier URI for OAUTH2 access token based authentication.
	SolClientSessionPropOauth2IssuerIdentifier = C.SOLCLIENT_SESSION_PROP_OAUTH2_ISSUER_IDENTIFIER
	// SolClientSessionPropOidcIDToken: The OIDC (OpenId Connect) ID Token. When authentication scheme ::SOLCLIENT_SESSION_PROP_AUTHENTICATION_SCHEME_OAUTH2 is used at least one of ::SOLCLIENT_SESSION_PROP_OAUTH2_ACCESS_TOKEN and ::SOLCLIENT_SESSION_PROP_OIDC_ID_TOKEN must be set.
	SolClientSessionPropOidcIDToken = C.SOLCLIENT_SESSION_PROP_OIDC_ID_TOKEN
	// SolClientSessionPropAuthenticationSchemeOauth2: OAUTH 2.0 authentication with a token.
	SolClientSessionPropAuthenticationSchemeOauth2 = C.SOLCLIENT_SESSION_PROP_AUTHENTICATION_SCHEME_OAUTH2
	// SolClientSessionPropSslValidateCertificateDate: This property indicates if the session connection should fail when a certificate with an invalid date (current date before valid date, or current date after expiry)  is received. This property only applies when ::SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE is enabled.<p>Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SSL_VALIDATE_CERTIFICATE_DATE.<p> See also @ref certificate-validation
	SolClientSessionPropSslValidateCertificateDate = C.SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE_DATE
	// SolClientSessionPropSslValidateCertificateHost: This property indicates if the session connection should fail when a certificate with an invalid host is received. When enabled, and connecting to a named host, the certificate Subject Alternative Name must contain a DNS entry that matches the host string. When enabled, and connecting to a host by IP address, the certificate Subject Alternative Name must contain an IP Address that matches.  If there is no Subject Alternate Name the certificate common name (CN) must match the named host.  This property only applies when ::SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE is enabled.<p> Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SSL_VALIDATE_CERTIFICATE_HOST.<p> See also @ref certificate-validation
	SolClientSessionPropSslValidateCertificateHost = C.SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE_HOST
	// SolClientSessionPropSslCipherSuites: This property specifies a comma separated list of the cipher suites. Allowed cipher suites are: 'ECDHE-RSA-AES256-GCM-SHA384', 'ECDHE-RSA-AES256-SHA384', 'ECDHE-RSA-AES256-SHA', 'AES256-GCM-SHA384', 'AES256-SHA256', 'AES256-SHA', 'ECDHE-RSA-DES-CBC3-SHA', 'DES-CBC3-SHA', 'ECDHE-RSA-AES128-GCM-SHA256', 'ECDHE-RSA-AES128-SHA256', 'ECDHE-RSA-AES128-SHA', 'AES128-GCM-SHA256', 'AES128-SHA256', 'AES128-SHA', 'RC4-SHA', 'RC4-MD5', 'TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384', 'TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384', 'TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA', 'TLS_RSA_WITH_AES_256_GCM_SHA384', 'TLS_RSA_WITH_AES_256_CBC_SHA256', 'TLS_RSA_WITH_AES_256_CBC_SHA', 'TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA', 'SSL_RSA_WITH_3DES_EDE_CBC_SHA', 'TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256', 'TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256', 'TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA', 'TLS_RSA_WITH_AES_128_GCM_SHA256', 'TLS_RSA_WITH_AES_128_CBC_SHA256', 'TLS_RSA_WITH_AES_128_CBC_SHA', 'SSL_RSA_WITH_RC4_128_SHA', 'SSL_RSA_WITH_RC4_128_MD5'. Default: ::SOLCLIENT_SESSION_PROP_DEFAULT_SSL_CIPHER_SUITES.
	SolClientSessionPropSslCipherSuites = C.SOLCLIENT_SESSION_PROP_SSL_CIPHER_SUITES
	// SolClientSessionPropSslTrustStoreDir: This property specifies the directory where the trusted certificates are. A maximum of 64 certificate files are allowed in the trust store directory. The maximum depth for the certificate chain verification that shall be allowed is 3.
	SolClientSessionPropSslTrustStoreDir = C.SOLCLIENT_SESSION_PROP_SSL_TRUST_STORE_DIR
	// SolClientSessionPropSslTrustedCommonNameList: <b>NOT RECOMMENDED</b>. Per RFC-6125 section 6.4.4, the Common Name (CN) of a certificate should not be used to validate the certificate.  Instead ::SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE_HOST (enabled by default) indicates that the server certificate must contain a Subject Alternate Name (SAN) that matches the host name (::SOLCLIENT_SESSION_PROP_HOST). <p>This property is only valid if ::SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE_HOST is disabled. This property specifies a comma separated list of acceptable common names in certificate validation. The number of common names specified by an applications is limited to 16. Leading and trailing whitespaces are considered to be part of the common names and are not ignored. If the application does not provide any common names, there is no common name verification.<p><b>NOTE:</b> When ::SOLCLIENT_SESSION_PROP_SSL_VALIDATE_CERTIFICATE_HOST is enabled (default), this property should be set to an empty list.  Failure to do so may cause the API to reject an otherwise valid certificate.
	SolClientSessionPropSslTrustedCommonNameList = C.SOLCLIENT_SESSION_PROP_SSL_TRUSTED_COMMON_NAME_LIST
	// SolClientSessionPropGdReconnectFailActionAutoRetry: Clear the publisher state and reconnect the publisher flow. Then republish all unacknowledged messages, this may cause duplication. The API then continues the reconnect process as usual.
	SolClientSessionPropGdReconnectFailActionAutoRetry = C.SOLCLIENT_SESSION_PROP_GD_RECONNECT_FAIL_ACTION_AUTO_RETRY
	// SolClientSessionPropGdReconnectFailActionDisconnect: Disconnect the session, even if SOLCLIENT_SESSION_PROP_RECONNECT_RETRIES is configured to a non-zero value. This is the legacy behavior. If the application attempts to manually reconnect the session, it is also responsible for unacknowledged messages. If the application chooses to resend those messages, there may be duplication. If the application chooses not to resend those messages there may be message loss. \n
	SolClientSessionPropGdReconnectFailActionDisconnect = C.SOLCLIENT_SESSION_PROP_GD_RECONNECT_FAIL_ACTION_DISCONNECT
	// SolClientSessionPropDefaultPort: The default value for the broker TCP port when compression is not in use (::SOLCLIENT_SESSION_PROP_COMPRESSION_LEVEL of zero).
	SolClientSessionPropDefaultPort = C.SOLCLIENT_SESSION_PROP_DEFAULT_PORT
	// SolClientSessionPropDefaultPortCompression: The default value for the broker TCP port when compression is in use (::SOLCLIENT_SESSION_PROP_COMPRESSION_LEVEL of non-zero).
	SolClientSessionPropDefaultPortCompression = C.SOLCLIENT_SESSION_PROP_DEFAULT_PORT_COMPRESSION
	// SolClientSessionPropDefaultPortSsl: The default value for the broker SSL port over TCP regardless of compression.
	SolClientSessionPropDefaultPortSsl = C.SOLCLIENT_SESSION_PROP_DEFAULT_PORT_SSL
	// SolClientSessionPropDefaultBufferSize: The default size (in bytes) of internal buffer for transmit buffering.
	SolClientSessionPropDefaultBufferSize = C.SOLCLIENT_SESSION_PROP_DEFAULT_BUFFER_SIZE
	// SolClientSessionPropDefaultBlockingWriteTimeoutMs: The default blocking write timeout (in milliseconds).
	SolClientSessionPropDefaultBlockingWriteTimeoutMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_BLOCKING_WRITE_TIMEOUT_MS
	// SolClientSessionPropDefaultConnectTimeoutMs: The default connect timeout (in milliseconds).
	SolClientSessionPropDefaultConnectTimeoutMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_CONNECT_TIMEOUT_MS
	// SolClientSessionPropDefaultSubconfirmTimeoutMs: The default subscription confirm (add or remove) timeout (in milliseconds).
	SolClientSessionPropDefaultSubconfirmTimeoutMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_SUBCONFIRM_TIMEOUT_MS
	// SolClientSessionPropDefaultSocketSendBufSize: Use 0 to set the socket send buffer size to the operating system default.
	SolClientSessionPropDefaultSocketSendBufSize = C.SOLCLIENT_SESSION_PROP_DEFAULT_SOCKET_SEND_BUF_SIZE
	// SolClientSessionPropDefaultSocketRcvBufSize: Use 0 to set the socket receive buffer size to the operating system default.
	SolClientSessionPropDefaultSocketRcvBufSize = C.SOLCLIENT_SESSION_PROP_DEFAULT_SOCKET_RCV_BUF_SIZE
	// SolClientSessionPropDefaultKeepAliveIntMs: The default amount of time (in milliseconds) to wait between sending out Keep-Alive messages.
	SolClientSessionPropDefaultKeepAliveIntMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_KEEP_ALIVE_INT_MS
	// SolClientSessionPropDefaultKeepAliveLimit: The default value for the number of consecutive Keep-Alive messages that can be sent without receiving a response before the connection is closed by the API.
	SolClientSessionPropDefaultKeepAliveLimit = C.SOLCLIENT_SESSION_PROP_DEFAULT_KEEP_ALIVE_LIMIT
	// SolClientSessionPropDefaultPubAckTimer: The default value for publisher acknowledgment timer (in milliseconds). When a published message is not acknowledged within the time specified for this timer, the API automatically retransmits the message. There is no limit on the number of retransmissions for any message. However, while the API is resending, applications can become flow controlled. The flow control behavior is controlled by ::SOLCLIENT_SESSION_PROP_SEND_BLOCKING and ::SOLCLIENT_SESSION_PROP_BLOCKING_WRITE_TIMEOUT_MS.
	SolClientSessionPropDefaultPubAckTimer = C.SOLCLIENT_SESSION_PROP_DEFAULT_PUB_ACK_TIMER
	// SolClientSessionPropDefaultPubWindowSize: The default Publisher Window size for Guaranteed messages. The Guaranteed Message Publish Window Size property limits the maximum number of messages that can be published before the API must receive an acknowledgment from the broker.
	SolClientSessionPropDefaultPubWindowSize = C.SOLCLIENT_SESSION_PROP_DEFAULT_PUB_WINDOW_SIZE
	// SolClientSessionPropDefaultSubscriberLocalPriority: The default subscriber priority for locally published messages.
	SolClientSessionPropDefaultSubscriberLocalPriority = C.SOLCLIENT_SESSION_PROP_DEFAULT_SUBSCRIBER_LOCAL_PRIORITY
	// SolClientSessionPropDefaultSubscriberNetworkPriority: The default subscriber priority for remotely published messages.
	SolClientSessionPropDefaultSubscriberNetworkPriority = C.SOLCLIENT_SESSION_PROP_DEFAULT_SUBSCRIBER_NETWORK_PRIORITY
	// SolClientSessionPropDefaultCompressionLevel: The default compression level (no compression).
	SolClientSessionPropDefaultCompressionLevel = C.SOLCLIENT_SESSION_PROP_DEFAULT_COMPRESSION_LEVEL
	// SolClientSessionPropDefaultConnectRetriesPerHost: The default number of connect retries per host. Zero means only try once when connecting.
	SolClientSessionPropDefaultConnectRetriesPerHost = C.SOLCLIENT_SESSION_PROP_DEFAULT_CONNECT_RETRIES_PER_HOST
	// SolClientSessionPropDefaultConnectRetries: The default number of connect retries. Zero means only try once when connecting.
	SolClientSessionPropDefaultConnectRetries = C.SOLCLIENT_SESSION_PROP_DEFAULT_CONNECT_RETRIES
	// SolClientSessionPropDefaultReconnectRetries: The default number of reconnect retries. Zero means no automatic connection retries after a Session goes down.
	SolClientSessionPropDefaultReconnectRetries = C.SOLCLIENT_SESSION_PROP_DEFAULT_RECONNECT_RETRIES
	// SolClientSessionPropDefaultReconnectRetryWaitMs: The default amount of time in (milliseconds) to wait before attempting a reconnect attempt.
	SolClientSessionPropDefaultReconnectRetryWaitMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_RECONNECT_RETRY_WAIT_MS
	// SolClientSessionPropDefaultProvisionTimeoutMs: The default amount of time (in milliseconds) to wait for a provision command.
	SolClientSessionPropDefaultProvisionTimeoutMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_PROVISION_TIMEOUT_MS
	// SolClientSessionPropDefaultModifypropTimeoutMs: The default amount of time (in milliseconds) to wait for session property modification.
	SolClientSessionPropDefaultModifypropTimeoutMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_MODIFYPROP_TIMEOUT_MS
	// SolClientSessionPropDefaultInitialReceiveBufferSize: The default value for ::SOLCLIENT_SESSION_PROP_INITIAL_RECEIVE_BUFFER_SIZE
	SolClientSessionPropDefaultInitialReceiveBufferSize = C.SOLCLIENT_SESSION_PROP_DEFAULT_INITIAL_RECEIVE_BUFFER_SIZE
	// SolClientSessionPropDefaultTransportProtocolDowngradeTimeoutMs: The default value for the Transport Protocol downgrade timeout in milliseconds.
	SolClientSessionPropDefaultTransportProtocolDowngradeTimeoutMs = C.SOLCLIENT_SESSION_PROP_DEFAULT_TRANSPORT_PROTOCOL_DOWNGRADE_TIMEOUT_MS
)
