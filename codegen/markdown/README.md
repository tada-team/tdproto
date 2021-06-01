## Structures

### ServerConfirm
Server confirmed client message



### ClientCallLeave
Leave call



### CallDevice
Call participant device

* **muted** (bool) — Device muted.
* **useragent** (string) — Device description.


### ClientCallTrickle
Send trickle candidate for webrtc connection



### ServerCallBuzz
Call buzzing



### Terms
Experimental translation fields for "team" entity renaming. Readonly



### ShortMessage
Short message based on chat message

* **chat** ( [JID](#JID) , readonly for clients) — Chat id.
* **chat_type** ( [ChatType](#ChatType) , readonly for clients) — Chat type.
* **created** (ISODateTimeString, readonly for clients) — Message creation datetime (set by server side) or sending datetime in future for draft messages.
* **from** ( [JID](#JID) , readonly for clients) — Sender contact id.
* **gentime** (int64, readonly for clients) — Object version.
* **is_archive** (bool, readonly for clients, omitempty) — This message is archive. True or null.
* **message_id** (string) — Message uid.
* **to** ( [JID](#JID) ) — Recipient id (group, task or contact).


### UploadShortMessage
Upload + ShortMessage



### ServerCallMuteall
All participants in call muted



### MessagePush
Push message over websockets. Readonly

* **chat** ( [JID](#JID) ) — Chat id.
* **click_action** (string) — Url opened on click.
* **created** (ISODateTimeString) — Message creation iso datetime.
* **icon_url** (string) — Absolute url to push icon.
* **message** (string) — Push body.
* **message_id** (string) — Message id.
* **sender** ( [JID](#JID) ) — Sender contact id.
* **subtitle** (string) — Push subtitle.
* **tag** (string) — Push tag (for join pushes).
* **team** (string) — Team uid.
* **title** (string) — Push title.


### ServerCallTalking
Someone talks in call



### Message
Chat message

* **chat** ( [JID](#JID) , readonly for clients) — Chat id.
* **chat_type** ( [ChatType](#ChatType) , readonly for clients) — Chat type.
* **content** ( [MessageContent](#MessageContent) ) — Message content struct.
* **created** (ISODateTimeString, readonly for clients) — Message creation datetime (set by server side) or sending datetime in future for draft messages.
* **_debug** (string, readonly for clients, omitempty) — Debug information, if any.
* **drafted** (ISODateTimeString, readonly for clients, omitempty) — Creation datetime for draft messages.
* **editable_until** (ISODateTimeString, readonly for clients, omitempty) — Author can change this message until date. Can be null.
* **edited** (ISODateTimeString, readonly for clients, omitempty) — ISODateTimeString of message modification or deletion.
* **from** ( [JID](#JID) , readonly for clients) — Sender contact id.
* **gentime** (int64, readonly for clients) — Object version.
* **has_previews** (bool, readonly for clients, omitempty) — Has link previews. True or null.
* **important** (bool, omitempty) — Importance flag.
* **is_archive** (bool, readonly for clients, omitempty) — This message is archive. True or null.
* **is_first** (bool, readonly for clients, omitempty) — This message is first in this chat. True or null.
* **is_last** (bool, readonly for clients, omitempty) — This message is first in this chat. True or null.
* **linked_messages** ( [Message](#Message) , list, omitempty) — Forwarded messages. Can be null. Also contains double of ReplyTo for backward compatibility.
* **links** ( [MessageLinks](#MessageLinks) , readonly for clients, omitempty) — External/internals links.
* **markup** ( [MarkupEntity](#MarkupEntity) , readonly for clients, list, omitempty) — Markup entities. Experimental.
* **message_id** (string) — Message uid.
* **nopreview** (bool, omitempty) — Disable link previews. True or null.
* **notice** (bool, readonly for clients, omitempty) — Has mention (@). True or null.
* **num** (int, readonly for clients, nullable, omitempty) — Index number of this message. Starts from 0. Null for deleted messages. Changes when any previous message wad deleted.
* **num_received** (int, readonly for clients, omitempty) — Unused yet.
* **prev** (string, readonly for clients, omitempty) — Previous message id in this chat. Uid or null.
* **push_text** (string, readonly for clients, omitempty) — Simple plaintext message representation.
* **reactions** ( [MessageReaction](#MessageReaction) , readonly for clients, list, omitempty) — Message reactions struct. Can be null.
* **received** (bool, readonly for clients, omitempty) — Message was seen by anybody in chat. True or null.
* **reply_to** ( [Message](#Message) , nullable, omitempty) — Message that was replied to, if any.
* **silently** (bool, readonly for clients, omitempty) — Message has no pushes and did not affect any counters.
* **to** ( [JID](#JID) ) — Recipient id (group, task or contact).
* **uploads** ( [Upload](#Upload) , list, omitempty) — Message uploads.


### ServerPanic
Critical server error



### DeletedTag
Delete tag message

* **uid** (string) — Tag id.


### Country
Country for phone numbers selection on login screen

* **code** (string) — Country code.
* **default** (bool, omitempty) — Selected by default.
* **name** (string) — Country name.
* **popular** (bool, omitempty) — Is popular, need to cache.


### ServerCallCheckFingerprint
Experimental function



### ClientPing
Empty message for checking server connection



### ServerCallReject
Call rejected



### Subtask
Link to sub/sup task

* **assignee** ( [JID](#JID) ) — Assignee contact id. Tasks only.
* **display_name** (string) — Title.
* **jid** ( [JID](#JID) ) — Task id.
* **num** (uint) — Task number in this team.
* **public** (bool, omitempty) — Can other team member see this task/group chat.
* **title** (string) — Task title. Generated from number and description.


### ServerTagUpdated
Tag created or changed



### IntegrationKind
Integration kind

* **description** (string) — Plugin description.
* **icon** (string) — Path to icon.
* **kind** (string) — Integration unique name.
* **template** ( [Integration](#Integration) ) — Integration template.
* **title** (string) — Plugin title.


### ColorRule
Set of rules to apply to tasks for coloring



### TeamShort
Short team representation. For invites, push notifications, etc. Readonly

* **icons** ( [IconData](#IconData) ) — Team icons.
* **name** (string) — Team name.
* **uid** (string) — Team id.


### ServerTime
Current server time



### ServerMessagePush
Push replacement for desktop application



### ServerTeamCounters
Counters form other teams



### MessageLinkPreview
Website title and description

* **description** (string, omitempty) — Website description.
* **title** (string) — Website title or og:title content.


### GroupMembership
Group chat membership status

* **can_remove** (bool, omitempty) — Can I remove this member.
* **jid** ( [JID](#JID) ) — Contact id.
* **status** ( [GroupStatus](#GroupStatus) ) — Status in group.


### ServerChatUpdated
Chat created or updated



### ServerCallLeave
Participant leave a call



### ServerRemindDeleted
Task or group remind deleted



### DeletedTeam
Team deletion message. Readonly

* **gentime** (int64) — Object version.
* **is_archive** (bool) — Team deleted.
* **uid** (string) — Team id.


### Session
Websocket session

* **addr** (string, omitempty) — IP address.
* **afk** (bool, omitempty) — Away from keyboard.
* **created** (ISODateTimeString) — Creation datetime.
* **is_mobile** (bool, omitempty) — Mobile.
* **lang** (string, omitempty) — Language code.
* **team** (string, omitempty) — Team id.
* **uid** (string) — Session id.
* **useragent** (string, omitempty) — User agent.


### Wallpaper
Chat wallpaper

* **key** (string) — Unique identifier.
* **name** (string) — Localized description.
* **url** (string) — Url to jpg or png.


### Upload
Uploaded media

* **animated** (bool, omitempty) — Is animated (images only).
* **blurhash** (string, omitempty) — Compact representation of a placeholder for an image (images only).
* **content_type** (string) — Content type.
* **created** (ISODateTimeString) — Uploaded at.
* **duration** (uint, omitempty) — Mediafile duration (for audio/video only).
* **type** ( [UploadMediaType](#UploadMediaType) ) — ?type=file,image,audio,video.
* **name** (string) — Filename.
* **pdf_version** ( [PdfVersion](#PdfVersion) , nullable, omitempty) — PDF version of file. Experimental.
* **preview** ( [UploadPreview](#UploadPreview) , nullable, omitempty) — Preview details.
* **processing** (bool, omitempty) — File still processing (video only).
* **size** (int) — Upload size in bytes.
* **uid** (string) — Upload id.
* **url** (string) — Absolute url.


### ClientMessageUpdated
Message created or changed



### ClientCallSound
Change mute state in call



### ServerCallState
Call participant number or parameters changed



### Integrations
Complete integrations data, as received from server

* **integrations** ( [Integration](#Integration) , list) — Currently existing integrations.
* **kinds** ( [IntegrationKind](#IntegrationKind) , list) — Types of integrations available for setup.


### MessageLink
Checked message links. In short: "Click here: {link.Pattern}" => "Click here: <a href='{link.Url}'>{link.Text}</a>"

* **nopreview** (bool, omitempty) — Website previews disabled.
* **pattern** (string) — Text fragment that should be replaced by link.
* **preview** ( [MessageLinkPreview](#MessageLinkPreview) , nullable, omitempty) — Optional preview info, for websites.
* **text** (string) — Text replacement.
* **uploads** ( [Upload](#Upload) , list, omitempty) — Optional upload info.
* **url** (string) — Internal or external link.
* **youtube_id** (string, omitempty) — Optional youtube movie id.


### TaskColor
Task color rules color



### Features
Server information. Readonly

* **afk_age** (int) — Max inactivity seconds.
* **allow_admin_mute** (bool, omitempty) — Wiki pages in chats. Experimental.
* **android_app** (string) — Link to Google Play.
* **app_schemes** (string, list) — Local applications urls.
* **app_title** (string) — Application title.
* **auth_2fa** (bool, omitempty) — Two-factor authentication (2FA) enabled.
* **auth_by_password** (bool, omitempty) — Password authentication enabled.
* **auth_by_qr_code** (bool, omitempty) — QR-code / link authentication enabled.
* **auth_by_sms** (bool, omitempty) — SMS authentication enabled.
* **background** (string, omitempty) — Background image url, if any.
* **build** (string) — Build/revision of server side.
* **calls** (bool) — Calls functions enabled.
* **calls_record** (bool) — Calls record enabled.
* **custom_server** (bool) — True for premise installation.
* **default_wallpaper** ( [Wallpaper](#Wallpaper) , nullable, omitempty) — Default wallpaper url for mobile apps, if any.
* **desktop_version** (string) — Desktop application version.
* **firebase_api_key** (string) — Firebase settings for web-push notifications.
* **firebase_app_id** (string) — Firebase settings for web-push notifications.
* **firebase_auth_domain** (string) — Firebase settings for web-push notifications.
* **firebase_database_url** (string) — Firebase settings for web-push notifications.
* **firebase_project_id** (string) — Firebase settings for web-push notifications.
* **firebase_sender_id** (string) — Firebase settings for web-push notifications.
* **firebase_storage_bucket** (string) — Firebase settings for web-push notifications.
* **free_registration** (bool) — Free registration allowed.
* **front_version** (string) — Webclient version.
* **host** (string) — Current host.
* **ice_servers** ( [ICEServer](#ICEServer) , list) — ICE servers for WebRTC.
* **ios_app** (string) — Link to AppStore.
* **installation_title** (string, omitempty) — Installation title, used on login screen.
* **installation_type** (string) — Name of installation.
* **is_testing** (bool) — Testing installation.
* **landing_url** (string, omitempty) — Landing page address, if any.
* **max_color_rule_description_length** (int) — Maximum length for ColorRule description.
* **max_group_title_length** (int) — Maximum chars for group chat name.
* **max_integration_comment_length** (int) — Maximum length for Integration comment.
* **max_linked_messages** (int) — Maximum number of forwarded messages.
* **max_message_length** (int) — Maximum chars for text message.
* **max_message_search_limit** (int) — Maximum search result.
* **max_message_uploads** (int) — Maximum number of message uploads.
* **max_mood_length** (int) — Maximum chars for mood in team.
* **max_participants_per_call** (int, omitempty) — Maximum number of participants per call.
* **max_role_length** (int) — Maximum chars for role in team.
* **max_section_length** (int) — Maximum length for project and contact's sections names.
* **max_tag_length** (int) — Maximum length for tags.
* **max_task_title_length** (int) — Maximum length for task title.
* **max_teams** (int) — Maximum teams for one account.
* **max_upload_mb** (int) — Maximum size of user's upload.
* **max_url_length** (int) — Maximum length for urls.
* **max_username_part_length** (int) — Maximum chars for: family_name, given_name, patronymic if any.
* **message_uploads** (bool) — Multiple message uploads.
* **metrika** (string) — Yandex metrika counter id.
* **min_app_version** (string) — Minimal application version required for this server. Used for breaking changes.
* **min_search_length** (int) — Minimal chars number for starting global search.
* **mobile_calls** (bool) — Calls functions enabled for mobile applications.
* **oauth_services** ( [OAuthService](#OAuthService) , list, omitempty) — External services.
* **only_one_device_per_call** (bool, omitempty) — Disallow call from multiply devices. Experimental.
* **readonly_groups** (bool) — Deprecated.
* **resend_timeout** (int) — Resend message in n seconds if no confirmation from server given.
* **safari_push_id** (string) — Safari push id for web-push notifications.
* **sentry_dsn_js** (string) — Frontend sentry.io settings.
* **server_drafts** (bool) — Message drafts saved on server.
* **single_group_teams** (bool) — Cross team communication. Experimental.
* **task_checklist** (bool) — Deprecated.
* **task_dashboard** (bool) — Deprecated.
* **task_messages** (bool) — Deprecated.
* **task_public** (bool) — Deprecated.
* **task_tags** (bool) — Deprecated.
* **terms** ( [Terms](#Terms) ) — Team entity naming. Experimental.
* **theme** (string) — Default UI theme.
* **userver** (string) — Static files server address.
* **wiki_pages** (bool) — Wiki pages in chats. Experimental.


### UserWithMe
Accouint data with extra information

* **devices** ( [PushDevice](#PushDevice) , list) — Registered push devices.
* **inviter** ( [JID](#JID) , omitempty) — Inviter id, if any.
* **teams** ( [Team](#Team) , list) — Available teams.
* **alt_send** (bool) — Use Ctrl/Cmd + Enter instead Enter.
* **always_send_pushes** (bool) — Send pushes even user is online.
* **asterisk_mention** (bool) — Use * as @ for mentions.
* **default_lang** (string, omitempty) — Default language code.
* **email** (string, omitempty) — Email for login.
* **family_name** (string, omitempty) — Family name.
* **given_name** (string, omitempty) — Given name.
* **munread_first** (bool) — Show unread chats in chat list first on mobiles.
* **patronymic** (string, omitempty) — Patronymic, if any.
* **phone** (string, omitempty) — Phone for login.
* **quiet_time_finish** (string, nullable) — Finish silently time (no pushes, no sounds).
* **quiet_time_start** (string, nullable) — Start silently time (no pushes, no sounds).
* **timezone** (string) — Timezone.
* **unread_first** (bool) — Show unread chats in chat list first.


### Unread
Unread message counters

* **chats** (uint) — Total chats with unread messages.
* **messages** (uint) — Total unread messages.
* **notice_messages** (uint) — Total unread messages with mentions.


### ServerLogin
Login from other device



### Team
Team

* **bad_profile** (bool, readonly for clients, omitempty) — My profile in this team isn't full.
* **changeable_statuses** ( [TeamStatus](#TeamStatus) , readonly for clients, list, omitempty) — What status I can set to other team members.
* **contacts** ( [Contact](#Contact) , readonly for clients, list, omitempty) — Team contacts. Used only for team creation.
* **default_task_deadline** (string, omitempty) — Default task deadline.
* **display_family_name_first** (bool, omitempty) — Family name should be first in display name.
* **gentime** (int64, readonly for clients) — Object version.
* **hide_archived_users** (bool, omitempty) — Don't show archived users by default.
* **icons** ( [IconData](#IconData) , readonly for clients) — Team icons.
* **is_archive** (bool, readonly for clients, omitempty) — Team deleted.
* **last_active** (bool, readonly for clients) — User last activity was in this team.
* **max_message_update_age** (int) — Max message update/deletion age, in seconds.
* **me** ( [Contact](#Contact) , readonly for clients) — My profile in this team.
* **name** (string) — Team name.
* **need_confirmation** (bool, readonly for clients) — Need confirmation after invite to this team.
* **single_group** ( [JID](#JID) , readonly for clients, omitempty) — For single group teams, jid of chat.
* **task_importance_max** (int, omitempty) — Maximum value of task importance. Default is 5.
* **task_importance_min** (int, omitempty) — Minimal value of task importance. Default is 1.
* **task_importance_rev** (bool, omitempty) — Bigger number = bigger importance. Default: lower number = bigger importance.
* **theme** ( [Theme](#Theme) , readonly for clients, nullable, omitempty) — Color theme, if any.
* **uid** (string, readonly for clients) — Team id.
* **unread** ( [TeamUnread](#TeamUnread) , readonly for clients, nullable) — Unread message counters.
* **uploads_size** (int64, readonly for clients, omitempty) — Total uploads size, bytes.
* **uploads_size_limit** (int64, readonly for clients, omitempty) — Maximum uploads size, bytes, if any.
* **use_patronymic** (bool, omitempty) — Patronymic in usernames for this team.
* **use_task_complexity** (bool, omitempty) — Use complexity field in task.
* **use_task_importance** (bool, omitempty) — Use importance field in task.
* **use_task_spent_time** (bool, omitempty) — Use spent time field in task.
* **use_task_urgency** (bool, omitempty) — Use urgency field in task.
* **user_fields** (string, readonly for clients, list) — Username fields ordering.


### ChatShort
Minimal chat representation

* **chat_type** ( [ChatType](#ChatType) ) — Chat type.
* **display_name** (string) — Title.
* **icons** ( [IconData](#IconData) , nullable) — Icon data.
* **jid** ( [JID](#JID) ) — Group/Task/Contact id.


### ServerChatDraft
Changed draft message in chan



### CallOnliner
Call participant

* **devices** ( [CallDevice](#CallDevice) , list) — Member devices, strictly one for now.
* **display_name** (string) — Contact name.
* **icon** (string) — Contact icon.
* **jid** ( [JID](#JID) ) — Contact id.
* **muted** (bool) — Microphone muted. Computed from devices muted states.
* **role** (string) — Contact role.


### TaskSort
Task sort type

* **key** ( [TaskSortKey](#TaskSortKey) ) — Field.
* **title** (string) — Sort title.


### ServerRemindFired
Task or group remind fired



### Contact
Contact

* **add_to_team_rights** (bool, omitempty) — Can contact add users to this team.
* **alt_send** (bool, nullable, omitempty) — Use Ctrl/Cmd + Enter instead Enter.
* **always_send_pushes** (bool, nullable, omitempty) — Send push notifications even contact is online.
* **asterisk_mention** (bool, nullable, omitempty) — Use * as @ for mentions.
* **auth_2fa_enabled** (bool, omitempty) — Two-factor authentication is configured and confirmed.
* **auth_2fa_status** (string, omitempty) — Two-factor authentication status.
* **botname** (string, omitempty) — Bot name. Empty for users.
* **can_add_to_group** (bool, omitempty) — Can I add this contact to group chats.
* **can_add_to_team** (bool, omitempty) — Can I add new members to this team.
* **can_call** (bool, omitempty) — Can I call to this contact.
* **can_create_group** (bool, omitempty) — Can I create group chats in this team.
* **can_create_task** (bool, omitempty) — Can I call create task for this contact.
* **can_delete** (bool, omitempty) — Can I remove this contact from team.
* **can_delete_any_message** (bool, omitempty) — Deprecated: use CanDeleteAnyMessage in chat object.
* **can_join_public_groups** (bool, omitempty) — Can I view/join public group in this team.
* **can_join_public_tasks** (bool, omitempty) — Can I view/join public tasks in this team.
* **can_manage_color_rules** (bool, omitempty) — Can I manage color rules in this team.
* **can_manage_integrations** (bool, omitempty) — Can I manage integrations in this team.
* **can_manage_sections** (bool, omitempty) — Can I manage sections in this team.
* **can_manage_tags** (bool, omitempty) — Can I manage tags in this team.
* **can_send_message** (bool, omitempty) — Can I send message to this contact.
* **cant_send_message_reason** (string, omitempty) — Why I can't send message to this chat (if can't).
* **changeable_fields** (string, list, omitempty) — Changeable fields.
* **contact_email** (string) — Contact email in this team.
* **contact_mshort_view** (bool, nullable, omitempty) — Short view in contact list in mobile app.
* **contact_phone** (string) — Contact phone in this team.
* **contact_short_view** (bool, nullable, omitempty) — Short view in contact list.
* **contact_show_archived** (bool, nullable, omitempty) — Show archived contacts in contact list.
* **custom_fields** ( [ContactCustomFields](#ContactCustomFields) , nullable, omitempty) — Extra contact fields.
* **debug_show_activity** (bool, nullable, omitempty) — Enable debug messages in UI.
* **default_lang** (string, nullable, omitempty) — Default language code.
* **display_name** (string) — Full name in chats.
* **dropall_enabled** (bool, nullable, omitempty) — Enable remove all messages experimental features.
* **family_name** (string, omitempty) — Family name.
* **given_name** (string, omitempty) — Given name.
* **group_mshort_view** (bool, nullable, omitempty) — Short view in group list in mobile app.
* **group_notifications_enabled** (bool, nullable, omitempty) — Push notifications for group chats.
* **group_short_view** (bool, nullable, omitempty) — Short view in group list.
* **icons** ( [IconData](#IconData) , nullable) — Icons data.
* **is_archive** (bool, omitempty) — Contact deleted.
* **jid** ( [JID](#JID) ) — Contact Id.
* **last_activity** (ISODateTimeString, omitempty) — Last activity in this team (iso datetime).
* **munread_first** (bool, nullable, omitempty) — Show unread chats first in feed in mobile app.
* **mood** (string, omitempty) — Mood in this team.
* **patronymic** (string, omitempty) — Patronymic, if any.
* **quiet_time_finish** (string, nullable, omitempty) — Quiet time finish.
* **quiet_time_start** (string, nullable, omitempty) — Quiet time start.
* **role** (string) — Role in this team.
* **sections** (string, list) — Section ids.
* **short_name** (string) — Short name in chats.
* **task_mshort_view** (bool, nullable, omitempty) — Short view in task list in mobile app.
* **task_notifications_enabled** (bool, nullable, omitempty) — Push notifications for task chats.
* **task_short_view** (bool, nullable, omitempty) — Short view in task list.
* **status** ( [TeamStatus](#TeamStatus) ) — Status in this team.
* **timezone** (string, nullable, omitempty) — Timezone, if any.
* **unread_first** (bool, nullable, omitempty) — Show unread chats first in feed.


### ServerCallRestart
Call restarted



### ServerTagDeleted
Tag deleted



### MessageContent
Chat message content

* **actor** ( [JID](#JID) , omitempty) — Change actor contact id (for "change" mediatype).
* **animated** (bool, omitempty) — Upload is animated image, if any. Deprecated: use Uploads instead.
* **blurhash** (string, omitempty) — Compact representation of a placeholder for an image. Deprecated: use Uploads instead.
* **comment** (string, omitempty) — Comment (for "audiomsg" mediatype).
* **duration** (uint, nullable, omitempty) — Upload duration, if any. Deprecated: use Uploads instead.
* **emails** (string, list, omitempty) — Emails list (for "contact" mediatype).
* **family_name** (string, omitempty) — Family name (for "contact" mediatype).
* **given_name** (string, omitempty) — Given name (for "contact" mediatype).
* **mediaURL** (string, omitempty) — Upload url, if any. Deprecated: use Uploads instead.
* **name** (string, omitempty) — Upload name, if any. Deprecated: use Uploads instead.
* **new** (string, nullable, omitempty) — Change new value (for "change" mediatype).
* **old** (string, nullable, omitempty) — Change old value (for "change" mediatype).
* **patronymic** (string, omitempty) — Patronymic name (for "contact" mediatype).
* **pdf_version** ( [PdfVersion](#PdfVersion) , nullable, omitempty) — Pdf version, if any.
* **phones** (string, list, omitempty) — Contact phones list (for "contact" mediatype).
* **preview2xURL** (string, omitempty) — Upload high resolution preview absolute url, if any. Deprecated: use Uploads instead.
* **previewHeight** (int, omitempty) — Upload preview height, in pixels, if any. Deprecated: use Uploads instead.
* **previewURL** (string, omitempty) — Upload preview absolute url, if any. Deprecated: use Uploads instead.
* **previewWidth** (int, omitempty) — Upload width, in pixels, if any. Deprecated: use Uploads instead.
* **processing** (bool, omitempty) — Upload still processing, if any. Deprecated: use Uploads instead.
* **size** (int, omitempty) — Upload size, if any. Deprecated: use Uploads instead.
* **stickerpack** (string, omitempty) — Stickerpack name (for "sticker" subtype).
* **subtype** ( [Mediasubtype](#Mediasubtype) , omitempty) — Message subtype, if any.
* **text** (string) — Text representation of message.
* **title** (string, omitempty) — Change title (for "change" mediatype).
* **type** ( [Mediatype](#Mediatype) ) — Message type.
* **upload** (string, omitempty) — Upload id, if any. Deprecated: use Uploads instead.


### ClientCallReject
Reject the call



### ServerSectionUpdated
Contact section or task project created or changed



### Remind
Remind

* **chat** ( [JID](#JID) ) — Chat id.
* **comment** (string, omitempty) — Comment, if any.
* **fire_at** (ISODateTimeString) — Activation time, iso.
* **uid** (string) — Remind id.


### ServerWarning
Something went wrong with client message



### ServerProcessing
Status of background operation



### ClientActivity
Change AFK (away from keyboard) status



### ServerChatLastread
Changed last read message in chat



### UploadPreview
Upload preview

* **height** (int) — Height in pixels.
* **url** (string) — Absolute url to image.
* **url_2x** (string) — Absolute url to high resolution image (retina).
* **width** (int) — Width in pixels.


### ServerCallAnswer
Call parameters



### ClientChatComposing
Typing or recording audiomessage



### User
Account data

* **alt_send** (bool) — Use Ctrl/Cmd + Enter instead Enter.
* **always_send_pushes** (bool) — Send pushes even user is online.
* **asterisk_mention** (bool) — Use * as @ for mentions.
* **default_lang** (string, omitempty) — Default language code.
* **email** (string, omitempty) — Email for login.
* **family_name** (string, omitempty) — Family name.
* **given_name** (string, omitempty) — Given name.
* **munread_first** (bool) — Show unread chats in chat list first on mobiles.
* **patronymic** (string, omitempty) — Patronymic, if any.
* **phone** (string, omitempty) — Phone for login.
* **quiet_time_finish** (string, nullable) — Finish silently time (no pushes, no sounds).
* **quiet_time_start** (string, nullable) — Start silently time (no pushes, no sounds).
* **timezone** (string) — Timezone.
* **unread_first** (bool) — Show unread chats in chat list first.


### TaskCounters
Tasks counters

* **jid** ( [JID](#JID) ) — Task jid.
* **num_unread** (uint, omitempty) — Unreads counter.
* **num_unread_notices** (uint, omitempty) — Mentions (@) counter.


### IntegrationForm
Integration form

* **api_key** ( [IntegrationField](#IntegrationField) , nullable, omitempty) — Api key field, if any.
* **url** ( [IntegrationField](#IntegrationField) , nullable, omitempty) — Url, if any.
* **webhook_url** ( [IntegrationField](#IntegrationField) , nullable, omitempty) — Webhook url, if any.


### ClientCallOffer
Start a call



### TaskItem
Task checklist item

* **can_toggle** (bool, omitempty) — Can I toggle this item.
* **checked** (bool, omitempty) — Item checked.
* **sort_ordering** (uint, omitempty) — Sort ordering.
* **subtask** ( [Subtask](#Subtask) , nullable, omitempty) — Link to subtask. Optional.
* **text** (string) — Text or "#{OtherTaskNumber}".
* **uid** (string, omitempty) — Id.


### IconData
Icon data. For icon generated from display name contains Letters + Color fields

* **blurhash** (string, omitempty) — Compact representation of a placeholder for an image (experimental).
* **color** (string, omitempty) — Icon background color (only for stub icon).
* **letters** (string, omitempty) — Letters (only for stub icon).
* **lg** ( [SingleIcon](#SingleIcon) ) — Large image.
* **sm** ( [SingleIcon](#SingleIcon) ) — Small icon.
* **stub** (string, omitempty) — Deprecated.


### TaskStatus
Custom task status

* **is_archive** (bool, omitempty) — Status not used anymore.
* **name** (string) — Status internal name.
* **sort_ordering** (uint) — Status sort ordering.
* **title** (string) — Status localized name.
* **uid** (string, omitempty) — Status id.


### TeamCounter
Unread message counters

* **uid** (string) — Team id.
* **unread** ( [TeamUnread](#TeamUnread) ) — Unread message counters.


### ServerRemindUpdated
Task/group remind created or changed



### ClientChatLastread
Last read message in chat changed



### ServerMessageReceived
Message received by someone in chat



### ServerChatComposing
Someone typing or recording audiomessage in chat



### ServerContactUpdated
Contact created or updated



### TaskTab
Task tab

* **filters** ( [TaskFilter](#TaskFilter) , list) — Filters inside tab.
* **hide_empty** (bool) — Disable this tab when it has no contents.
* **key** ( [TaskTabKey](#TaskTabKey) ) — Tab name.
* **pagination** (bool) — Enable pagination.
* **show_counter** (bool) — Show unread badge.
* **sort** ( [TaskSort](#TaskSort) , list) — Sort available in tab.
* **title** (string) — Tab title.
* **unread_tasks** ( [TaskCounters](#TaskCounters) , list) — Unread tasks with jid and counters.


### PdfVersion
PDF preview of mediafile. Experimental

* **text_preview** (string, omitempty) — First string of text content.
* **url** (string) — Absolute url.


### CallEvent
Audio call information

* **audiorecord** (bool) — Call record enabled.
* **finish** (ISODateTimeString, nullable) — Call finish.
* **onliners** ( [CallOnliner](#CallOnliner) , list, omitempty) — Call participants.
* **start** (ISODateTimeString, nullable) — Call start.


### Integration
Integration for concrete chat

* **comment** (string) — Comment, if any.
* **created** (ISODateTimeString, omitempty) — Creation datetime, iso.
* **enabled** (bool) — Integration enabled.
* **form** ( [IntegrationForm](#IntegrationForm) ) — Integration form.
* **group** ( [JID](#JID) ) — Chat id.
* **help** (string, omitempty) — Full description.
* **kind** (string) — Unique integration name.
* **uid** (string, omitempty) — Id.


### OAuthService
OAuth service

* **name** (string) — Integration title.
* **url** (string) — Redirect url.


### ContactCustomFields
Extra contact fields



### MarkupEntity
Markup entity. Experimental

* **childs** ( [MarkupEntity](#MarkupEntity) , list, omitempty) — List of internal markup entities.
* **cl** (int) — Close marker offset.
* **cllen** (int, omitempty) — Close marker length.
* **op** (int) — Open marker offset.
* **oplen** (int, omitempty) — Open marker length.
* **repl** (string, omitempty) — Text replacement.
* **time** (string, omitempty) — Time, for Time type.
* **typ** ( [MarkupType](#MarkupType) ) — Marker type.
* **url** (string, omitempty) — Url, for Link type.


### ServerChatDeleted
Chat deleted



### TaskFilter
Task filter

* **field** ( [TaskFilterKey](#TaskFilterKey) ) — Task filter field.
* **title** (string) — Filter title.


### ServerUploadUpdated
Upload object created or changed



### ServerTeamDeleted
Team archived



### ServerCallSound
Mute/unmute call participant



### ClientConfirm
Client confirmed server message



### ServerUiSettings
Part of UI settings changed



### WikiPage
Wiki page. Experimental

* **editor** ( [JID](#JID) ) — Last editor contact id.
* **gentime** (int64) — Object version.
* **text** (string) — Page text.
* **updated** (ISODateTimeString) — Update time.


### ICEServer
Interactive Connectivity Establishment Server for WEB Rtc connection. Readonly

* **urls** (string) — URls.


### ServerCallBuzzcancel
Call cancelled on buzzing



### ServerTeamUpdated
Team created or changed



### DeletedRemind
Remind deleted message

* **uid** (string) — Remind id.


### SingleIcon
Small or large icon

* **height** (int) — Icon height, in pixels.
* **url** (string) — absolute url to icon.
* **width** (int) — Icon width, in pixels.


### MessageReaction
Message emoji reaction

* **counter** (int) — Number of reactions.
* **details** ( [MessageReactionDetail](#MessageReactionDetail) , list) — Details.
* **name** (string) — Emoji.


### ClientMessageDeleted
Message deleted



### Chat
Chat (direct, group, task) representation

* **assignee** ( [JID](#JID) , omitempty) — Assignee contact id. Tasks only.
* **autocleanup_age** (int, nullable, omitempty) — Delete messages in this chat in seconds. Experimental function.
* **base_gentime** (int64, omitempty) — Base fields (not related to concrete participant) version.
* **can_add_member** (bool, omitempty) — Can I add member to this group chat.
* **can_call** (bool, omitempty) — Can I call to this chat.
* **can_change_member_status** (bool, omitempty) — Can I change member status in this group chat.
* **can_change_settings** (bool, omitempty) — deprecated: use changeable fields.
* **can_delete** (bool, omitempty) — Can I delete this chat.
* **can_delete_any_message** (bool, omitempty) — Can I delete any message in this chat.
* **can_join** (bool, omitempty) — Can I join to this public group/task.
* **can_remove_member** (bool, omitempty) — Can I remove member from this group chat.
* **can_send_message** (bool, omitempty) — Can I send message to this chat.
* **can_set_important_any_message** (bool, omitempty) — Can I change Important flag in any message in this chat.
* **cant_send_message_reason** (string, omitempty) — Why I can't send message to this chat (if can't).
* **changeable_fields** (string, list, omitempty) — List of editable fields.
* **chat_type** ( [ChatType](#ChatType) ) — Chat type.
* **collapsed** (bool, omitempty) — Description collapsed. Used for tasks only.
* **color_index** (uint16, nullable, omitempty) — Custom color index from table of colors. Tasks only.
* **complexity** (int, nullable, omitempty) — Task complexity, number.
* **counters_enabled** (bool, omitempty) — Include unread messages to counters.
* **created** (ISODateTimeString) — Creation date, iso datetime.
* **deadline** (ISODateTimeString, omitempty) — Task deadline in iso format, if any.
* **deadline_expired** (bool, omitempty) — Is task deadline expired.
* **default_for_all** (bool, omitempty) — Any new team member will be added to this group chat.
* **description** (string, omitempty) — Group or task description.
* **display_name** (string) — Title.
* **done** (ISODateTimeString, omitempty) — Task done date in iso format, if any.
* **done_reason** (string, omitempty) — Task done reason, if any.
* **draft** (string, omitempty) — Last message draft, if any.
* **draft_num** (int64, omitempty) — Last message draft version , if any.
* **feed** (bool, omitempty) — Present in feed (main screen).
* **gentime** (int64) — Chat fields related to concrete participant) version.
* **hidden** (bool, omitempty) — Hidden chat.
* **icons** ( [IconData](#IconData) , nullable) — Icons info.
* **importance** (int, nullable, omitempty) — Task importance, if available in team.
* **items** ( [TaskItem](#TaskItem) , list, omitempty) — Checklist items. Task only.
* **jid** ( [JID](#JID) ) — Group/Task/Contact id.
* **last_activity** (ISODateTimeString, omitempty) — Date of the last message sent even if it was deleted.
* **last_message** ( [Message](#Message) , nullable, omitempty) — Last message object.
* **last_read_message_id** (string, omitempty) — Last read message id, if any.
* **linked_messages** (interface{}, list, omitempty) — Used for "Create task from messages...".
* **links** ( [MessageLinks](#MessageLinks) , omitempty) — Links in description.
* **markup** ( [MarkupEntity](#MarkupEntity) , readonly for clients, list, omitempty) — Markup entities for description field. Experimental.
* **members** ( [GroupMembership](#GroupMembership) , list, omitempty) — Group chat members.
* **notifications_enabled** (bool, omitempty) — Push notifications enabled.
* **num** (uint, omitempty) — Task number in this team.
* **num_checked_items** (uint, nullable, omitempty) — Checked items in checklist. Tasks only.
* **num_importants** (int, omitempty) — Number of important messages.
* **num_items** (uint, nullable, omitempty) — Items in checklist. Tasks only.
* **num_members** (uint, nullable, omitempty) — Non-archive participants number.
* **num_unread** (uint, omitempty) — Unread counter.
* **num_unread_notices** (uint, omitempty) — Mentions (@) counter.
* **observers** ( [JID](#JID) , list, omitempty) — Task followers id's. TODO: rename to "followers".
* **owner** ( [JID](#JID) , omitempty) — Task creator.
* **parents** ( [Subtask](#Subtask) , list, omitempty) — Parent tasks.
* **pinned** (bool, omitempty) — Is chat pinned on top.
* **pinned_message** ( [Message](#Message) , nullable, omitempty) — Pinned message for this chat.
* **pinned_sort_ordering** (int, omitempty) — Sort ordering for pinned chat.
* **public** (bool, omitempty) — Can other team member see this task/group chat.
* **readonly_for_members** (bool, omitempty) — Readonly for non-admins group chat (Like Channels in Telegram bug switchable).
* **section** (string, omitempty) — Project / section id, if any.
* **spent_time** (int, nullable, omitempty) — Task spent time, number.
* **status** ( [GroupStatus](#GroupStatus) , nullable, omitempty) — My status in group chat.
* **tabs** ( [TaskTabKey](#TaskTabKey) , list, omitempty) — Tab names.
* **tags** (string, list, omitempty) — Task tags list, if any.
* **task_status** (string, omitempty) — Task status. May be custom.
* **title** (string, omitempty) — Task title. Generated from number and description.
* **uploads** ( [Upload](#Upload) , list, omitempty) — Upload uids for request, upload objects for response.
* **urgency** (int, nullable, omitempty) — Task urgency, if available in team.


### Theme
Color theme

* **AppAccentColor** (string) — Deprecated.
* **AppPrimaryColor** (string) — Deprecated.
* **BgColor** (string) — Web colors.
* **brand** (string) — App colors.


### Tag
Task tag

* **name** (string) — Tag name.
* **uid** (string) — Tag id.


### ServerDebug
Debug message



### IntegrationField
Integration form field

* **label** (string) — Label.
* **readonly** (bool) — Is field readonly.
* **value** (string) — Current value.


### MessageReactionDetail
Message reaction detail

* **created** (ISODateTimeString) — When reaction added, iso datetime.
* **name** (string) — Reaction emoji.
* **sender** ( [JID](#JID) ) — Reaction author.


### ContactShort
Short contact representation

* **display_name** (string) — Full name in chats.
* **icons** ( [IconData](#IconData) , nullable) — Icons data.
* **jid** ( [JID](#JID) ) — Contact Id.
* **short_name** (string) — Short name in chats.


### ServerOnline
Online team members and current active calls



### ClientCallMuteAll
Mute all other call participants



### ServerSectionDeleted
Contact section or task project deleted



### clientCallBuzzParams
Call buzzing

* **jid** ( [JID](#JID) ) — Chat or contact id.
* **members** ( [JID](#JID) , list) — List of call participants.


### DeletedChat
Minimal chat representation for deletion

* **chat_type** ( [ChatType](#ChatType) ) — Chat type.
* **gentime** (int64) — Chat fields (related to concrete participant) version.
* **is_archive** (bool) — Archive flag. Always true for this structure.
* **jid** ( [JID](#JID) ) — Group/Task/Contact id.


### ClientCallBuzzCancel
Call buzzing cancelled



### ServerMessageUpdated
Chat message created, updated or deleted


## Events

### "client.activity"

Event structure: [ClientActivity](#ClientActivity)

Change AFK (away from keyboard) status

```
{
	"confirm_id": "75a406625c58",
	"event": "client.activity",
	"params": {
		"afk": "BOOL"
	}
}
```
### "client.call.buzz"

Event structure: [ClientCallBuzz](#ClientCallBuzz)



```
{
	"confirm_id": "64977b08d763",
	"event": "client.call.buzz",
	"params": {
		"jid": "JID"
	}
}
```
### "client.call.buzzcancel"

Event structure: [ClientCallBuzzCancel](#ClientCallBuzzCancel)

Call buzzing cancelled

```
{
	"confirm_id": "8c52201ff7ed",
	"event": "client.call.buzzcancel",
	"params": {
		"jid": "JID"
	}
}
```
### "client.call.leave"

Event structure: [ClientCallLeave](#ClientCallLeave)

Leave call

```
{
	"confirm_id": "f5b6d29013c3",
	"event": "client.call.leave",
	"params": {
		"jid": "JID",
		"reason": "STRING"
	}
}
```
### "client.call.muteall"

Event structure: [ClientCallMuteAll](#ClientCallMuteAll)

Mute all other call participants

```
EVENT MISSING EXAMPLE
```
### "client.call.offer"

Event structure: [ClientCallOffer](#ClientCallOffer)

Start a call

```
{
	"confirm_id": "b45fdc034116",
	"event": "client.call.offer",
	"params": {
		"jid": "JID",
		"muted": "BOOL",
		"sdp": "STRING",
		"trickle": "BOOL"
	}
}
```
### "client.call.reject"

Event structure: [ClientCallReject](#ClientCallReject)

Reject the call

```
{
	"confirm_id": "55e8cc25d534",
	"event": "client.call.reject",
	"params": {
		"jid": "JID"
	}
}
```
### "client.call.sound"

Event structure: [ClientCallSound](#ClientCallSound)

Change mute state in call

```
{
	"confirm_id": "4a24b770a659",
	"event": "client.call.sound",
	"params": {
		"jid": "JID",
		"muted": "BOOL"
	}
}
```
### "client.call.trickle"

Event structure: [ClientCallTrickle](#ClientCallTrickle)

Send trickle candidate for webrtc connection

```
{
	"confirm_id": "5bde78b37316",
	"event": "client.call.trickle",
	"params": {
		"candidate": "STRING",
		"jid": "JID",
		"sdp_mid": "STRING",
		"sdp_mline_index": "INT"
	}
}
```
### "client.chat.composing"

Event structure: [ClientChatComposing](#ClientChatComposing)

Typing or recording audiomessage

```
{
	"confirm_id": "2bd5afaf39af",
	"event": "client.chat.composing",
	"params": {
		"jid": "JID"
	}
}
```
### "client.chat.lastread"

Event structure: [ClientChatLastread](#ClientChatLastread)

Last read message in chat changed

```
{
	"confirm_id": "8561d892f3d8",
	"event": "client.chat.lastread",
	"params": {
		"jid": "JID",
		"last_read_message_id": "STRING"
	}
}
```
### "client.confirm"

Event structure: [ClientConfirm](#ClientConfirm)

Client confirmed server message

```
{
	"event": "client.confirm",
	"params": {
		"confirm_id": "str"
	}
}
```
### "client.message.deleted"

Event structure: [ClientMessageDeleted](#ClientMessageDeleted)

Message deleted

```
{
	"confirm_id": "cd778785149a",
	"event": "client.message.deleted",
	"params": {
		"message_id": "STRING"
	}
}
```
### "client.message.updated"

Event structure: [ClientMessageUpdated](#ClientMessageUpdated)

Message created or changed

```
{
	"confirm_id": "2694a2864526",
	"event": "client.message.updated",
	"params": {
		"comment": "STRING",
		"content": {
			"text": "...",
			"type": "plain"
		},
		"important": "BOOL",
		"linked_messages": [
			"STRING"
		],
		"message_id": "STRING",
		"nopreview": "BOOL",
		"to": "JID"
	}
}
```
### "client.ping"

Event structure: [ClientPing](#ClientPing)

Empty message for checking server connection

```
{
	"confirm_id": "8aad294579b8",
	"event": "client.ping"
}
```
### "server.call.answer"

Event structure: [ServerCallAnswer](#ServerCallAnswer)

Call parameters

```
{
    "event": "server.call.answer",
    "params": {
        "candidates": [
            {
                "candidate": "candidate:2 1 tcp 1019216383 95.213.181.86 0 typ host tcptype active",
                "sdpMLineIndex": 0
            }
        ],
        "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2",
        "jsep": {
            "sdp": ".....SDP....",
            "type": "answer"
        },
        "onliners": [
            {
                "devices": [
                    {
                        "browser": null,
                        "muted": true
                    },
                    {
                        "browser": null,
                        "muted": true
                    }
                ],
                "jid": "d-ef53637c-f44c-4f49-9ffb-05436eb995ce",
                "muted": true
            }
        ]
    }
}
```
### "server.call.buzz"

Event structure: [ServerCallBuzz](#ServerCallBuzz)

Call buzzing

```
{
    "event": "server.call.buzz",
    "params": {
        "display_name": "имя фамилия",
        "icons": {
            "lg": {
                "height": 512,
                "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                "width": 512
            },
            "sm": {
                "height": 256,
                "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                "width": 256
            }
        },
        "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2",
        "team": "6aefcf3b-e61c-4f35-8b5b-9d762a6a6cf9"
    }
}
```
### "server.call.buzzcancel"

Event structure: [ServerCallBuzzcancel](#ServerCallBuzzcancel)

Call cancelled on buzzing

```
{
    "event": "server.call.buzzcancel",
    "params": {
        "jid": "d-0bdfbbf5-abfa-4ed2-9f98-991d5bb80127",
        "team": "848cc926-3048-44b3-a9ba-3195a394351d"
    }
}
```
### "server.call.checkfingerprint"

Event structure: [ServerCallCheckFingerprint](#ServerCallCheckFingerprint)

Experimental function

```
EVENT MISSING EXAMPLE
```
### "server.call.leave"

Event structure: [ServerCallLeave](#ServerCallLeave)

Participant leave a call

```
EVENT MISSING EXAMPLE
```
### "server.call.muteall"

Event structure: [ServerCallMuteall](#ServerCallMuteall)

All participants in call muted

```
EVENT MISSING EXAMPLE
```
### "server.call.reject"

Event structure: [ServerCallReject](#ServerCallReject)

Call rejected

```
EVENT MISSING EXAMPLE
```
### "server.call.restart"

Event structure: [ServerCallRestart](#ServerCallRestart)

Call restarted

```
EVENT MISSING EXAMPLE
```
### "server.call.sound"

Event structure: [ServerCallSound](#ServerCallSound)

Mute/unmute call participant

```
EVENT MISSING EXAMPLE
```
### "server.call.state"

Event structure: [ServerCallState](#ServerCallState)

Call participant number or parameters changed

```
{
    "event": "server.call.state",
    "params": {
        "audiorecord": true,
        "buzz": false,
        "finish": null,
        "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2",
        "onliners": [
            {
                "devices": [
                    {
                        "muted": false,
                        "useragent": "Chrome"
                    }
                ],
                "jid": "d-d6de9748-0bcd-4145-b1ce-3d7f41c1c26f",
                "muted": false
            }
        ],
        "start": "2019-09-27T01:01:35.264095Z"
    }
}
```
### "server.call.talking"

Event structure: [ServerCallTalking](#ServerCallTalking)

Someone talks in call

```
{
    "event": "server.call.talking",
    "params": {
        "_level": null,
        "actor": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
        "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
        "talking": true
    }
}
```
### "server.chat.composing"

Event structure: [ServerChatComposing](#ServerChatComposing)

Someone typing or recording audiomessage in chat

```
{
    "event": "server.chat.composing",
    "params": {
        "actor": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
        "composing": true,
        "is_audio": false,
        "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e"
    }
}
```
### "server.chat.deleted"

Event structure: [ServerChatDeleted](#ServerChatDeleted)

Chat deleted

```
{
    "event": "server.chat.deleted",
    "params": {
        "badge": 449,
        "chats": [
            {
                "chat_type": "group",
                "gentime": 1569546094920474368,
                "is_archive": true,
                "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2"
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 1,
                "messages": 31
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 1,
                "messages": 1
            }
        }
    }
}
```
### "server.chat.draft"

Event structure: [ServerChatDraft](#ServerChatDraft)

Changed draft message in chan

```
EVENT MISSING EXAMPLE
```
### "server.chat.lastread"

Event structure: [ServerChatLastread](#ServerChatLastread)

Changed last read message in chat

```
{
    "event": "server.chat.lastread",
    "params": {
        "badge": 0,
        "chats": [
            {
                "chat_type": "direct",
                "gentime": 1569546084300075149,
                "jid": "d-2a9d8c43-41d1-479c-9c8c-f029f799a724",
                "last_read_message_id": "f4f1f09c-f978-4b54-a321-b4c82604010d",
                "num_unread": 0,
                "num_unread_notices": 0
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 0,
                "messages": 0
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 0,
                "messages": 0
            }
        }
    }
}
```
### "server.chat.updated"

Event structure: [ServerChatUpdated](#ServerChatUpdated)

Chat created or updated

```
{
    "event": "server.chat.updated",
    "params": {
        "badge": 449,
        "chats": [
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "group",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:06.105362Z",
                "default_for_all": true,
                "description": "Сюда автоматически попадают все участники команды",
                "display_name": "Общий чат",
                "gentime": 1568764808134709248,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9E%D1%87/256.png"
                },
                "jid": "g-82d6c0d4-8921-4519-abb3-02f17373112f",
                "notifications_enabled": true,
                "num_members": 32,
                "public": true,
                "status": "member"
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 1,
                "messages": 31
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 1,
                "messages": 1
            }
        }
    }
}
```
### "server.confirm"

Event structure: [ServerConfirm](#ServerConfirm)

Server confirmed client message

```
{
	"event": "server.confirm",
	"params": {
		"confirm_id": "b8b2ccd6-35a6-408f-a591-c696a9f9e73e"
	}
}
```
### "server.contact.updated"

Event structure: [ServerContactUpdated](#ServerContactUpdated)

Contact created or updated

```
{
    "event": "server.contact.updated",
    "params": {
        "contacts": [
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000031",
                "display_name": "Контакт №32",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/32/256.png"
                },
                "jid": "d-a792832a-d9f6-4bbc-b545-f2800e329c8c",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            }
        ]
    }
}
```
### "server.debug"

Event structure: [ServerDebug](#ServerDebug)

Debug message

```
EVENT MISSING EXAMPLE
```
### "server.login"

Event structure: [ServerLogin](#ServerLogin)

Login from other device

```
{
    "event": "server.login",
    "params": {
        "device_name": "(unknown device)"
    }
}
```
### "server.message.push"

Event structure: [ServerMessagePush](#ServerMessagePush)

Push replacement for desktop application

```
EVENT MISSING EXAMPLE
```
### "server.message.received"

Event structure: [ServerMessageReceived](#ServerMessageReceived)

Message received by someone in chat

```
{
    "event": "server.message.received",
    "params": {
        "messages": [
            {
                "chat": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74",
                "message_id": "e0fca4c2-0d17-4317-96cd-703c2ca781e4",
                "num_received": 0,
                "received": false
            }
        ]
    }
}
```
### "server.message.updated"

Event structure: [ServerMessageUpdated](#ServerMessageUpdated)

Chat message created, updated or deleted

```
{
    "event": "server.message.updated",
    "params": {
        "badge": 51,
        "chat_counters": [
            {
                "chat_type": "task",
                "gentime": 1569546093662061959,
                "jid": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74",
                "last_read_message_id": null,
                "num_unread": 0,
                "num_unread_notices": 0
            }
        ],
        "delayed": false,
        "messages": [
            {
                "chat": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74",
                "chat_type": "task",
                "content": {
                    "actor": "d-0b1b2362-fb69-4dd6-8bfb-90c89517168d",
                    "text": "Создана задача для @Контакт_№394: Тест",
                    "title": "Создана задача для @Контакт_№394: Тест",
                    "type": "change"
                },
                "created": "2019-09-27T01:01:33.661466Z",
                "editable_until": "2019-09-28T01:01:33.661466Z",
                "from": "d-a2d94e8e-74aa-4b0b-b55d-0b4fc1ce07f4",
                "gentime": 1569546093662061959,
                "is_first": true,
                "is_last": true,
                "message_id": "e0fca4c2-0d17-4317-96cd-703c2ca781e4",
                "to": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74"
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 1,
                "messages": 51
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 0,
                "messages": 0
            }
        }
    }
}
```
### "server.online"

Event structure: [ServerOnline](#ServerOnline)

Online team members and current active calls

```
{
    "event": "server.online",
    "params": {
        "calls": [
            {
                "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2"
            }
        ],
        "contacts": [
            {
                "afk": false,
                "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                "mobile": false
            }
        ]
    }
}
```
### "server.panic"

Event structure: [ServerPanic](#ServerPanic)

Critical server error

```
{
	"event": "server.panic",
	"params": {
		"code": "INVALID_TEAM",
		"debug": null
	}
}
```
### "server.processing"

Event structure: [ServerProcessing](#ServerProcessing)

Status of background operation

```
EVENT MISSING EXAMPLE
```
### "server.remind.deleted"

Event structure: [ServerRemindDeleted](#ServerRemindDeleted)

Task or group remind deleted

```
{
    "event": "server.remind.deleted",
    "params": {
        "reminds": [
            {
                "uid": "22018199-c3ae-4a9c-829e-985e975eb62a"
            }
        ]
    }
}
```
### "server.remind.fired"

Event structure: [ServerRemindFired](#ServerRemindFired)

Task or group remind fired

```
{
    "event": "server.remind.fired",
    "params": {
        "reminds": [
            {
                "chat": "t-2c527748-2e00-440e-a625-58121d42fdad",
                "fire_at": "2019-09-27T11:01:03.428564Z",
                "uid": "22018199-c3ae-4a9c-829e-985e975eb62a"
            }
        ]
    }
}
```
### "server.remind.updated"

Event structure: [ServerRemindUpdated](#ServerRemindUpdated)

Task/group remind created or changed

```
{
    "event": "server.remind.updated",
    "params": {
        "reminds": [
            {
                "chat": "t-2c527748-2e00-440e-a625-58121d42fdad",
                "fire_at": "2019-09-27T11:01:03.428564Z",
                "uid": "22018199-c3ae-4a9c-829e-985e975eb62a"
            }
        ]
    }
}
```
### "server.roster"

Event structure: [ServerRoster](#ServerRoster)



```
{
    "event": "server.roster",
    "params": {
        "badge": 449,
        "contact_sections": [
            {
                "gentime": 1568764805859444992,
                "name": "Боты",
                "sort_ordering": 99,
                "uid": "3a8b23f2-25b4-4acb-afab-590303bedd50"
            }
        ],
        "contacts": [
            {
                "botname": "systembot",
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "support@tada.team",
                "contact_phone": "",
                "display_name": "tada.teamBot",
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                        "width": 256
                    }
                },
                "jid": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                "last_activity": "2019-09-27T01:01:32.354349Z",
                "role": "Виртуальный ассистент",
                "section": "3a8b23f2-25b4-4acb-afab-590303bedd50",
                "sections": [
                    "3a8b23f2-25b4-4acb-afab-590303bedd50"
                ],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000000",
                "display_name": "Контакт №1",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-181b9859-ea3f-4e26-85c9-9e7ee7aaaa9a",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "admin"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000009",
                "display_name": "Контакт №10",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/10/256.png"
                },
                "jid": "d-a87e168a-78df-49aa-9abb-9b5409eb7de5",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000010",
                "display_name": "Контакт №11",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/11/256.png"
                },
                "jid": "d-8bb2ed5e-5cf4-457c-9a7b-070d7e4116a9",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000011",
                "display_name": "Контакт №12",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/12/256.png"
                },
                "jid": "d-5049173a-93dc-4904-8796-6807b6f07bf1",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000012",
                "display_name": "Контакт №13",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/13/256.png"
                },
                "jid": "d-26879ae0-a4b9-4523-8875-4e26c4787d27",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000013",
                "display_name": "Контакт №14",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/14/256.png"
                },
                "jid": "d-84e3cb48-598b-4ffd-918e-0662ebeb3469",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000014",
                "display_name": "Контакт №15",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/15/256.png"
                },
                "jid": "d-6496b27c-3251-4683-896a-a7f2191b7a33",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000015",
                "display_name": "Контакт №16",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/16/256.png"
                },
                "jid": "d-c540dd6e-4da3-4b65-8e89-bd35554978c2",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000016",
                "display_name": "Контакт №17",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/17/256.png"
                },
                "jid": "d-951775af-15fd-4e5f-bb13-099c9afd6586",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000017",
                "display_name": "Контакт №18",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/18/256.png"
                },
                "jid": "d-ce8117ae-867d-400d-b7c8-e7df6a2ecee7",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000018",
                "display_name": "Контакт №19",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/19/256.png"
                },
                "jid": "d-1d60b554-5da7-4f17-a078-eb77168081a8",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "alt_send": false,
                "can_add_to_group": true,
                "can_create_task": true,
                "can_delete": true,
                "changeable_fields": [
                    "alt_send",
                    "contact_email",
                    "contact_mshort_view",
                    "contact_phone",
                    "contact_short_view",
                    "contact_show_archived",
                    "debug_show_activity",
                    "default_lang",
                    "family_name",
                    "given_name",
                    "group_mshort_view",
                    "group_notifications_enabled",
                    "group_short_view",
                    "icons",
                    "phone",
                    "role",
                    "task_mshort_view",
                    "task_notifications_enabled",
                    "task_short_view",
                    "unread_first"
                ],
                "contact_email": "",
                "contact_mshort_view": false,
                "contact_phone": "+75550000001",
                "contact_short_view": false,
                "contact_show_archived": false,
                "debug_show_activity": false,
                "default_lang": null,
                "display_name": "Контакт №2",
                "family_name": "№2",
                "given_name": "Контакт",
                "group_mshort_view": false,
                "group_notifications_enabled": true,
                "group_short_view": false,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member",
                "task_mshort_view": false,
                "task_notifications_enabled": true,
                "task_short_view": false,
                "unread_first": false
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000019",
                "display_name": "Контакт №20",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/20/256.png"
                },
                "jid": "d-13b47a4f-4473-4a8f-82ef-00b00ce5c02f",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000020",
                "display_name": "Контакт №21",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/21/256.png"
                },
                "jid": "d-14ba8969-e3a9-4daf-b6c3-bed4ea931e06",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000021",
                "display_name": "Контакт №22",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/22/256.png"
                },
                "jid": "d-571d1843-96d1-497b-a53c-2cc03d815194",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000022",
                "display_name": "Контакт №23",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/23/256.png"
                },
                "jid": "d-c15b9ad4-b03b-4dc1-aeaa-725e09bedbcf",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000023",
                "display_name": "Контакт №24",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/24/256.png"
                },
                "jid": "d-0f902404-96c4-4c5e-8385-d33a46fd5566",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000024",
                "display_name": "Контакт №25",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/25/256.png"
                },
                "jid": "d-b54d8559-4d94-46e6-b1ff-673a55e5cf5d",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000025",
                "display_name": "Контакт №26",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/26/256.png"
                },
                "jid": "d-d62ca7c6-83b5-4853-bfc5-10e8c2e68b13",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000026",
                "display_name": "Контакт №27",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/27/256.png"
                },
                "jid": "d-7589ac63-4c0d-4685-9bf5-63b83cdcc3ab",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000027",
                "display_name": "Контакт №28",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/28/256.png"
                },
                "jid": "d-10e10caf-6127-4649-b2d6-5ad1395d32a5",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000028",
                "display_name": "Контакт №29",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/29/256.png"
                },
                "jid": "d-1b20a9aa-ecf5-49fd-8477-0998d42e1ade",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000002",
                "display_name": "Контакт №3",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-ee3d7c07-fe0f-4cbd-bf3a-b8e72f0ff1c6",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000029",
                "display_name": "Контакт №30",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/30/256.png"
                },
                "jid": "d-bc7550f7-ae4b-41de-82d3-d4b5b5b7f600",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000030",
                "display_name": "Контакт №31",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/31/256.png"
                },
                "jid": "d-bfe1fa23-c32b-48ea-8fed-d2b6f583fb77",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000031",
                "display_name": "Контакт №32",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/32/256.png"
                },
                "jid": "d-a792832a-d9f6-4bbc-b545-f2800e329c8c",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000003",
                "display_name": "Контакт №4",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-c094f142-1b54-4fae-8128-9b41504b56a9",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000004",
                "display_name": "Контакт №5",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-53e898d1-3d7b-4cc4-9ff6-ca034b5709cb",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000005",
                "display_name": "Контакт №6",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-af8d0ae3-ff07-4c49-94d0-7d531f972b0e",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000006",
                "display_name": "Контакт №7",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-f73f2847-4fcb-4f7e-bc11-8d2aaf21cf00",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000007",
                "display_name": "Контакт №8",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-2bde58fa-3b2b-4c10-8041-b0a65e6d74b4",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000008",
                "display_name": "Контакт №9",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-425ef763-2162-4dcf-8e0e-5ce0615518d8",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "botname": "notebot",
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "",
                "display_name": "Мои заметки",
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot256.png",
                        "width": 256
                    }
                },
                "jid": "d-6d02dd61-8c07-4d11-a1f5-4584569d7b6d",
                "last_activity": "2019-09-27T01:01:32.354092Z",
                "role": "Блокнот",
                "section": "3a8b23f2-25b4-4acb-afab-590303bedd50",
                "sections": [
                    "3a8b23f2-25b4-4acb-afab-590303bedd50"
                ],
                "status": "member"
            }
        ],
        "devices": [],
        "direct_chats": [
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:42.318781Z",
                "display_name": "Контакт №32",
                "gentime": 1568764842326248192,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/32/256.png"
                },
                "jid": "d-a792832a-d9f6-4bbc-b545-f2800e329c8c",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:41.097659Z",
                "display_name": "Контакт №31",
                "gentime": 1568764841104200192,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/31/256.png"
                },
                "jid": "d-bfe1fa23-c32b-48ea-8fed-d2b6f583fb77",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:38.428242Z",
                "display_name": "Контакт №30",
                "gentime": 1568764838432389632,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/30/256.png"
                },
                "jid": "d-bc7550f7-ae4b-41de-82d3-d4b5b5b7f600",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:36.676433Z",
                "display_name": "Контакт №29",
                "gentime": 1568764836680841984,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/29/256.png"
                },
                "jid": "d-1b20a9aa-ecf5-49fd-8477-0998d42e1ade",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:34.945690Z",
                "display_name": "Контакт №28",
                "gentime": 1568764834949238528,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/28/256.png"
                },
                "jid": "d-10e10caf-6127-4649-b2d6-5ad1395d32a5",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:33.326639Z",
                "display_name": "Контакт №27",
                "gentime": 1568764833332984320,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/27/256.png"
                },
                "jid": "d-7589ac63-4c0d-4685-9bf5-63b83cdcc3ab",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:31.700784Z",
                "display_name": "Контакт №26",
                "gentime": 1568764831705233664,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/26/256.png"
                },
                "jid": "d-d62ca7c6-83b5-4853-bfc5-10e8c2e68b13",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:30.429840Z",
                "display_name": "Контакт №25",
                "gentime": 1568764830435267328,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/25/256.png"
                },
                "jid": "d-b54d8559-4d94-46e6-b1ff-673a55e5cf5d",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:28.982264Z",
                "display_name": "Контакт №24",
                "gentime": 1568764828988044800,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/24/256.png"
                },
                "jid": "d-0f902404-96c4-4c5e-8385-d33a46fd5566",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:27.429493Z",
                "display_name": "Контакт №23",
                "gentime": 1568764827437857536,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/23/256.png"
                },
                "jid": "d-c15b9ad4-b03b-4dc1-aeaa-725e09bedbcf",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:26.171104Z",
                "display_name": "Контакт №22",
                "gentime": 1568764826174592768,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/22/256.png"
                },
                "jid": "d-571d1843-96d1-497b-a53c-2cc03d815194",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:25.083140Z",
                "display_name": "Контакт №21",
                "gentime": 1568764825085393152,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/21/256.png"
                },
                "jid": "d-14ba8969-e3a9-4daf-b6c3-bed4ea931e06",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:23.970206Z",
                "display_name": "Контакт №20",
                "gentime": 1568764823973247744,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/20/256.png"
                },
                "jid": "d-13b47a4f-4473-4a8f-82ef-00b00ce5c02f",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:23.013194Z",
                "display_name": "Контакт №19",
                "gentime": 1568764823017283072,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/19/256.png"
                },
                "jid": "d-1d60b554-5da7-4f17-a078-eb77168081a8",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:21.974534Z",
                "display_name": "Контакт №18",
                "gentime": 1568764821976428800,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/18/256.png"
                },
                "jid": "d-ce8117ae-867d-400d-b7c8-e7df6a2ecee7",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:21.113294Z",
                "display_name": "Контакт №17",
                "gentime": 1568764821115754496,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/17/256.png"
                },
                "jid": "d-951775af-15fd-4e5f-bb13-099c9afd6586",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:20.254806Z",
                "display_name": "Контакт №16",
                "gentime": 1568764820257041664,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/16/256.png"
                },
                "jid": "d-c540dd6e-4da3-4b65-8e89-bd35554978c2",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:19.406973Z",
                "display_name": "Контакт №15",
                "gentime": 1568764819409007360,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/15/256.png"
                },
                "jid": "d-6496b27c-3251-4683-896a-a7f2191b7a33",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:18.430679Z",
                "display_name": "Контакт №14",
                "gentime": 1568764818432459008,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/14/256.png"
                },
                "jid": "d-84e3cb48-598b-4ffd-918e-0662ebeb3469",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:17.577868Z",
                "display_name": "Контакт №13",
                "gentime": 1568764817580576512,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/13/256.png"
                },
                "jid": "d-26879ae0-a4b9-4523-8875-4e26c4787d27",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:16.423089Z",
                "display_name": "Контакт №12",
                "gentime": 1568764816433606656,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/12/256.png"
                },
                "jid": "d-5049173a-93dc-4904-8796-6807b6f07bf1",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:15.370702Z",
                "display_name": "Контакт №11",
                "gentime": 1568764815376827136,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/11/256.png"
                },
                "jid": "d-8bb2ed5e-5cf4-457c-9a7b-070d7e4116a9",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:14.302001Z",
                "display_name": "Контакт №10",
                "gentime": 1568764814322221056,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/10/256.png"
                },
                "jid": "d-a87e168a-78df-49aa-9abb-9b5409eb7de5",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:13.306725Z",
                "display_name": "Контакт №9",
                "gentime": 1568764813310813440,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-425ef763-2162-4dcf-8e0e-5ce0615518d8",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:12.378013Z",
                "display_name": "Контакт №8",
                "gentime": 1568764812385766400,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-2bde58fa-3b2b-4c10-8041-b0a65e6d74b4",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:11.451771Z",
                "display_name": "Контакт №7",
                "gentime": 1568764811455242240,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-f73f2847-4fcb-4f7e-bc11-8d2aaf21cf00",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:10.605227Z",
                "display_name": "Контакт №6",
                "gentime": 1568764810616535296,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-af8d0ae3-ff07-4c49-94d0-7d531f972b0e",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:09.803211Z",
                "display_name": "Контакт №5",
                "gentime": 1568764809805365248,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-53e898d1-3d7b-4cc4-9ff6-ca034b5709cb",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:09.082753Z",
                "display_name": "Контакт №4",
                "gentime": 1568764809084504320,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-c094f142-1b54-4fae-8128-9b41504b56a9",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:08.302823Z",
                "display_name": "Контакт №3",
                "gentime": 1568764808307773184,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-ee3d7c07-fe0f-4cbd-bf3a-b8e72f0ff1c6",
                "notifications_enabled": true
            },
            {
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:07.568804Z",
                "display_name": "tada.teamBot",
                "gentime": 1568764899109919304,
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                        "width": 256
                    }
                },
                "jid": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                "last_message": {
                    "chat": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                    "chat_type": "direct",
                    "content": {
                        "text": "Новый участник команды: @Контакт_№32",
                        "type": "plain"
                    },
                    "created": "2019-09-18T00:01:39.103766Z",
                    "from": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                    "gentime": 1568764899109919304,
                    "is_last": true,
                    "links": [
                        {
                            "_notify_ids": [
                                34
                            ],
                            "_status": 1,
                            "pattern": "@Контакт_№32",
                            "text": "@Контакт №32",
                            "url": "otv://d-a792832a-d9f6-4bbc-b545-f2800e329c8c"
                        }
                    ],
                    "message_id": "6610b580-1d97-46da-902c-36fe03db9392",
                    "prev": "27980080-6d7c-4389-b2df-958bac09c801",
                    "to": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e"
                },
                "notifications_enabled": true,
                "num_unread": 31
            },
            {
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:07.517104Z",
                "display_name": "Мои заметки",
                "gentime": 1568764807520538368,
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot256.png",
                        "width": 256
                    }
                },
                "jid": "d-6d02dd61-8c07-4d11-a1f5-4584569d7b6d",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:07.435409Z",
                "display_name": "Контакт №1",
                "gentime": 1568764807438561792,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-181b9859-ea3f-4e26-85c9-9e7ee7aaaa9a",
                "notifications_enabled": true
            }
        ],
        "group_sections": [],
        "groups": [
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "group",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:06.105362Z",
                "default_for_all": true,
                "description": "Сюда автоматически попадают все участники команды",
                "display_name": "Общий чат",
                "gentime": 1568764808134709248,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9E%D1%87/256.png"
                },
                "jid": "g-82d6c0d4-8921-4519-abb3-02f17373112f",
                "notifications_enabled": true,
                "num_members": 32,
                "public": true,
                "status": "member"
            }
        ],
        "task_chats": [
            {
                "assignee": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                "can_call": true,
                "can_send_message": true,
                "chat_type": "task",
                "counters_enabled": true,
                "created": "2019-09-18T00:23:46.506612Z",
                "display_name": "1. Загрузить фото профиля",
                "gentime": 1568766227276560299,
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                        "width": 256
                    }
                },
                "jid": "t-2185340c-2956-4433-a4be-c42c73a5f226",
                "last_message": {
                    "chat": "t-2185340c-2956-4433-a4be-c42c73a5f226",
                    "chat_type": "task",
                    "content": {
                        "actor": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                        "text": "Created task for @Контакт_№2: Загрузить фото профиля\nВ мобильном приложении настройки вашего профиля находятся слева, в боковом меню. В веб-версии — в правом верхнем углу.",
                        "title": "Created task for @Контакт_№2: Загрузить фото профиля\nВ мобильном приложении настройки вашего профиля находятся слева, в боковом меню. В веб-версии — в правом верхнем углу.",
                        "type": "change"
                    },
                    "created": "2019-09-18T00:23:47.276536Z",
                    "from": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                    "gentime": 1568766227276560299,
                    "is_first": true,
                    "is_last": true,
                    "links": [
                        {
                            "_notify_ids": [
                                4
                            ],
                            "_status": 1,
                            "pattern": "@Контакт_№2",
                            "text": "@Контакт №2",
                            "url": "otv://d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e"
                        }
                    ],
                    "message_id": "977f228b-2aa5-4b8b-a4f2-0ff36116aacc",
                    "notice": true,
                    "to": "t-2185340c-2956-4433-a4be-c42c73a5f226"
                },
                "notifications_enabled": true,
                "num": 1,
                "num_members": 2,
                "num_unread": 1,
                "num_unread_notices": 1,
                "observers": [],
                "owner": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                "tabs": [
                    "active",
                    "in"
                ],
                "task_status": "new",
                "title": "Загрузить фото профиля"
            }
        ],
        "task_sections": [],
        "task_tabs": [
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": false,
                "key": "active",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Все"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": true,
                "key": "expired",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Истёк срок"
            },
            {
                "filters": [
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": false,
                "key": "in",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Входящие"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": false,
                "key": "out",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Исходящие"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": true,
                "key": "copy",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Наблюдаю"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": true,
                "key": "done",
                "pagination": true,
                "show_counter": false,
                "sort": [
                    {
                        "key": "done",
                        "title": "Дата завершения"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Готово"
            }
        ]
    }
}
```
### "server.section.deleted"

Event structure: [ServerSectionDeleted](#ServerSectionDeleted)

Contact section or task project deleted

```
{
    "event": "server.section.deleted",
    "params": {
        "chat_type": "direct",
        "gentime": 1568786364815431424,
        "sections": [
            {
                "uid": "8daeb4a2-4863-4953-bcf1-c58b84dc5272"
            }
        ]
    }
}
```
### "server.section.updated"

Event structure: [ServerSectionUpdated](#ServerSectionUpdated)

Contact section or task project created or changed

```
{
    "event": "server.section.updated",
    "params": {
        "chat_type": "group",
        "gentime": 1569546093241191168,
        "sections": []
    }
}
```
### "server.tag.deleted"

Event structure: [ServerTagDeleted](#ServerTagDeleted)

Tag deleted

```
EVENT MISSING EXAMPLE
```
### "server.tag.updated"

Event structure: [ServerTagUpdated](#ServerTagUpdated)

Tag created or changed

```
EVENT MISSING EXAMPLE
```
### "server.team.counters"

Event structure: [ServerTeamCounters](#ServerTeamCounters)

Counters form other teams

```
{
    "event": "server.team.counters",
    "params": {
        "teams": [
            {
                "uid": "b3e92407-41d0-456a-9b1d-3d6b79bfa7d3",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 161
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "8b05a648-d43a-4d17-9a0e-c2a119453494",
                "unread": {
                    "direct": {
                        "chats": 0,
                        "messages": 0
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "6aefcf3b-e61c-4f35-8b5b-9d762a6a6cf9",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 3
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 0,
                        "messages": 0
                    }
                }
            },
            {
                "uid": "5e854825-d1ac-4b28-a389-f0e8cc3bd675",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 1
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 0,
                        "messages": 0
                    }
                }
            },
            {
                "uid": "180b9b6d-ecb0-47aa-8712-33dbd9b4a642",
                "unread": {
                    "direct": {
                        "chats": 0,
                        "messages": 0
                    },
                    "group": {
                        "chats": 1,
                        "messages": 58
                    },
                    "task": {
                        "chats": 0,
                        "messages": 0
                    }
                }
            },
            {
                "uid": "99cabb3b-d4fd-464d-b33d-144b67e3c1df",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 127
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "523e7bdc-e73f-4150-b6b9-656d6178ba51",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 63
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "c9d8a896-a2b6-40a1-869e-2ecc0ef2436b",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 31
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            }
        ]
    }
}
```
### "server.team.deleted"

Event structure: [ServerTeamDeleted](#ServerTeamDeleted)

Team archived

```
{
    "event": "server.team.deleted",
    "params": {
        "teams": [
            {
                "is_archive": true,
                "uid": "c9d8a896-a2b6-40a1-869e-2ecc0ef2436b"
            }
        ]
    }
}
```
### "server.team.updated"

Event structure: [ServerTeamUpdated](#ServerTeamUpdated)

Team created or changed

```
{
    "event": "server.team.updated",
    "params": {
        "teams": [
            {
                "can_manage_sections": true,
                "default_task_deadline": "18:00:00",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/00/256.png"
                },
                "last_active": false,
                "max_message_update_age": 86400,
                "me": {
                    "alt_send": false,
                    "can_add_to_group": true,
                    "can_create_task": true,
                    "can_delete": true,
                    "changeable_fields": [
                        "alt_send",
                        "contact_email",
                        "contact_mshort_view",
                        "contact_phone",
                        "contact_short_view",
                        "contact_show_archived",
                        "debug_show_activity",
                        "default_lang",
                        "family_name",
                        "given_name",
                        "group_mshort_view",
                        "group_notifications_enabled",
                        "group_short_view",
                        "icons",
                        "phone",
                        "role",
                        "task_mshort_view",
                        "task_notifications_enabled",
                        "task_short_view",
                        "unread_first"
                    ],
                    "contact_email": "",
                    "contact_mshort_view": false,
                    "contact_phone": "+75550000001",
                    "contact_short_view": false,
                    "contact_show_archived": false,
                    "debug_show_activity": false,
                    "default_lang": null,
                    "display_name": "Контакт №2",
                    "family_name": "№2",
                    "given_name": "Контакт",
                    "group_mshort_view": false,
                    "group_notifications_enabled": true,
                    "group_short_view": false,
                    "icons": {
                        "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                    },
                    "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                    "last_activity": null,
                    "role": "",
                    "sections": [],
                    "status": "member",
                    "task_mshort_view": false,
                    "task_notifications_enabled": true,
                    "task_short_view": false,
                    "unread_first": false
                },
                "name": "T-32 2019-09-18 00:00:05.603182+00:00",
                "need_confirmation": true,
                "uid": "c9d8a896-a2b6-40a1-869e-2ecc0ef2436b",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 31
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            }
        ]
    }
}
```
### "server.time"

Event structure: [ServerTime](#ServerTime)

Current server time

```
{
    "event": "server.time",
    "params": {
        "time": "2019-09-27T01:01:33.216665Z"
    }
}
```
### "server.uisettings"

Event structure: [ServerUiSettings](#ServerUiSettings)

Part of UI settings changed

```
{
	"last-tab": 91238475
}
```
### "server.upload.updated"

Event structure: [ServerUploadUpdated](#ServerUploadUpdated)

Upload object created or changed

```
EVENT MISSING EXAMPLE
```
### "server.warning"

Event structure: [ServerWarning](#ServerWarning)

Something went wrong with client message

```
{
	"event": "server.warning",
	"params": {
		"message": "unknown event: client.deeeeeaddddbeeeeeef",
		"orig": {
			"confirm_id": "c6280a82ed1c",
			"event": "client.deeeeeaddddbeeeeeef"
		}
	}
}
```