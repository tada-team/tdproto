## Structures

### ActiveUserDailyStat
MISSING CLASS DOCUMENTATION

* **call_seconds_total** (int, nullable, omitempty) — DOCUMENTATION MISSING.

* **calls_count** (int, nullable, omitempty) — DOCUMENTATION MISSING.

* **day** (time.Time) — DOCUMENTATION MISSING.

* **family_name** (string, nullable, omitempty) — DOCUMENTATION MISSING.

* **given_name** (string, nullable, omitempty) — DOCUMENTATION MISSING.

* **messages_count** (int, nullable, omitempty) — DOCUMENTATION MISSING.

* **patronymic** (string, nullable, omitempty) — DOCUMENTATION MISSING.

* **phone** (string, nullable, omitempty) — DOCUMENTATION MISSING.

* **user_id** (int) — DOCUMENTATION MISSING.



### AnyEvent
MISSING CLASS DOCUMENTATION



### BaseEvent
MISSING CLASS DOCUMENTATION

* **confirm_id** (string, omitempty) — DOCUMENTATION MISSING.

* **event** (string) — DOCUMENTATION MISSING.



### ButtonColors
MISSING CLASS DOCUMENTATION

* **brand_active** (string) — DOCUMENTATION MISSING.

* **brand_disable** (string) — DOCUMENTATION MISSING.

* **brand_static** (string) — DOCUMENTATION MISSING.

* **simple_active** (string) — DOCUMENTATION MISSING.

* **simple_disable** (string) — DOCUMENTATION MISSING.

* **simple_static** (string) — DOCUMENTATION MISSING.



### CallDevice
Call participant device

* **muted** (bool) — Device muted.

* **useragent** (string) — Device description.



### CallEvent
Audio call information

* **audiorecord** (bool) — Call record enabled.

* **finish** (string, nullable) — Call finish, iso date.

* **onliners** ( [CallOnliner](#CallOnliner) , list) — Call participants.

* **start** (string, nullable) — Call start, iso date.



### CallOnliner
Call participant

* **devices** ( [CallDevice](#CallDevice) , list) — Member devices, strictly one for now.

* **display_name** (string) — Contact name.

* **icon** (string) — Contact icon.

* **jid** ( [JID](#JID) ) — Contact id.

* **muted** (bool) — Microphone muted. Computed from devices muted states.

* **role** (string) — Contact role.



### Chat
Chat (direct, group, task) representation

* **assignee** ( [JID](#JID) , nullable, omitempty) — Assignee contact id. Tasks only.

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

* **deadline** (string, omitempty) — Task deadline in iso format, if any.

* **deadline_expired** (bool, omitempty) — Is task deadline expired.

* **default_for_all** (bool, omitempty) — Any new team member will be added to this group chat.

* **description** (string, omitempty) — Group or task description.

* **display_name** (string) — Title.

* **done** (string, omitempty) — Task done date in iso format, if any.

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

* **observers** ( [JID](#JID) , nullable, list, omitempty) — Task followers id's. TODO: rename to "followers".

* **owner** ( [JID](#JID) , nullable, omitempty) — Task creator.

* **parents** ( [Subtask](#Subtask) , list, omitempty) — Parent tasks.

* **pinned** (bool, omitempty) — Is chat pinned on top.

* **pinned_message** ( [Message](#Message) , nullable, omitempty) — Pinned message for this chat.

* **pinned_sort_ordering** (int, omitempty) — Sort ordering for pinned chat.

* **public** (bool, omitempty) — Can other team member see this task/group chat.

* **readonly_for_members** (bool, omitempty) — Readonly for non-admins group chat (Like Channels in Telegram bug switchable).

* **section** (string, omitempty) — Project / section id, if any.

* **spent_time** (int, nullable, omitempty) — Task spent time, number.

* **status** ( [GroupStatus](#GroupStatus) , nullable, omitempty) — My status in group chat.

* **tabs** ( [TaskTabKey](#TaskTabKey) , nullable, list, omitempty) — Tab names.

* **tags** (string, list, omitempty) — Task tags list, if any.

* **task_status** (string, omitempty) — Task status. May be custom.

* **title** (string, omitempty) — Task title. Generated from number and description.

* **uploads** ( [Upload](#Upload) , list, omitempty) — Upload uids for request, upload objects for response.

* **urgency** (int, nullable, omitempty) — Task urgency, if available in team.



### ChatCounters
MISSING CLASS DOCUMENTATION

* **chat_type** ( [ChatType](#ChatType) ) — DOCUMENTATION MISSING.

* **gentime** (int64) — DOCUMENTATION MISSING.

* **jid** ( [JID](#JID) ) — DOCUMENTATION MISSING.

* **last_activity** (ISODateTimeString, omitempty) — DOCUMENTATION MISSING.

* **last_read_message_id** (string, nullable) — DOCUMENTATION MISSING.

* **num_unread** (uint) — DOCUMENTATION MISSING.

* **num_unread_notices** (uint) — DOCUMENTATION MISSING.



### ChatMessages
MISSING CLASS DOCUMENTATION

* **messages** ( [Message](#Message) , list) — DOCUMENTATION MISSING.



### ChatShort
Minimal chat representation

* **chat_type** ( [ChatType](#ChatType) ) — Chat type.

* **display_name** (string) — Title.

* **icons** ( [IconData](#IconData) , nullable) — Icon data.

* **jid** ( [JID](#JID) ) — Group/Task/Contact id.



### ClientActivity
Change AFK (away from keyboard) status

* **params** ( [clientActivityParams](#clientActivityParams) ) — DOCUMENTATION MISSING.



### ClientCallBuzzCancel
Call buzzing cancelled

* **params** ( [clientCallBuzzCancelParams](#clientCallBuzzCancelParams) ) — DOCUMENTATION MISSING.



### ClientCallLeave
Leave call

* **params** ( [clientCallLeaveParams](#clientCallLeaveParams) ) — DOCUMENTATION MISSING.



### ClientCallMuteAll
Mute all other call participants

* **params** ( [clientCallMuteAllParams](#clientCallMuteAllParams) ) — DOCUMENTATION MISSING.



### ClientCallOffer
Start a call

* **params** ( [clientCallOfferParams](#clientCallOfferParams) ) — DOCUMENTATION MISSING.



### ClientCallReject
Reject the call

* **params** ( [clientCallRejectParams](#clientCallRejectParams) ) — DOCUMENTATION MISSING.



### ClientCallSound
Change mute state in call

* **params** ( [clientCallSoundParams](#clientCallSoundParams) ) — DOCUMENTATION MISSING.



### ClientCallTrickle
Send trickle candidate for webrtc connection

* **params** ( [clientCallTrickleParams](#clientCallTrickleParams) ) — DOCUMENTATION MISSING.



### ClientChatComposing
Typing or recording audiomessage

* **params** ( [clientChatComposingParams](#clientChatComposingParams) ) — DOCUMENTATION MISSING.



### ClientChatLastread
Last read message in chat changed

* **params** ( [clientChatLastreadParams](#clientChatLastreadParams) ) — DOCUMENTATION MISSING.



### ClientConfirm
Client confirmed server message

* **params** ( [clientConfirmParams](#clientConfirmParams) ) — DOCUMENTATION MISSING.



### ClientMessageDeleted
Message deleted

* **params** ( [clientMessageDeletedParams](#clientMessageDeletedParams) ) — DOCUMENTATION MISSING.



### ClientMessageUpdated
Message created or changed

* **params** ( [ClientMessageUpdatedParams](#ClientMessageUpdatedParams) ) — DOCUMENTATION MISSING.



### ClientMessageUpdatedParams
MISSING CLASS DOCUMENTATION

* **comment** (string, omitempty) — Deprecated.

* **content** ( [MessageContent](#MessageContent) ) — Message content. Required.

* **important** (bool, omitempty) — Important flag. Not required. Default: false.

* **linked_messages** (string, list, omitempty) — Forwarded messages (previously was for reply too). Not required.

* **message_id** (string, omitempty) — Uid created by client. Recommended.

* **nopreview** (bool, omitempty) — Disable links preview generation. Not required. Default: false.

* **old_style_attachment** (bool, omitempty) — Backward compatibility mode.

* **reply_to** (string, omitempty) — Replied to message id. Not required.

* **to** ( [JID](#JID) ) — Chat, task or contact jid. Required.

* **uploads** (string, list, omitempty) — Message attachments.



### ClientPing
Empty message for checking server connection



### ColorRule
Set of rules to apply to tasks for coloring

* **color_index** (uint16) — DOCUMENTATION MISSING.

* **description** (string, omitempty) — DOCUMENTATION MISSING.

* **priority** (int) — DOCUMENTATION MISSING.

* **section** (string, omitempty) — DOCUMENTATION MISSING.

* **section_enabled** (bool, nullable, omitempty) — DOCUMENTATION MISSING.

* **tags** (string, list, omitempty) — DOCUMENTATION MISSING.

* **tags_enabled** (bool, nullable, omitempty) — DOCUMENTATION MISSING.

* **task_importance** (int, nullable, omitempty) — DOCUMENTATION MISSING.

* **task_importance_enabled** (bool, nullable, omitempty) — DOCUMENTATION MISSING.

* **task_status** (string, omitempty) — DOCUMENTATION MISSING.

* **task_urgency** (int, nullable, omitempty) — DOCUMENTATION MISSING.

* **task_urgency_enabled** (bool, nullable, omitempty) — DOCUMENTATION MISSING.

* **uid** (string) — DOCUMENTATION MISSING.



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

* **changeable_fields** (string, nullable, list, omitempty) — Changeable fields.

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

* **family_name** (string, nullable, omitempty) — Family name.

* **given_name** (string, nullable, omitempty) — Given name.

* **group_mshort_view** (bool, nullable, omitempty) — Short view in group list in mobile app.

* **group_notifications_enabled** (bool, nullable, omitempty) — Push notifications for group chats.

* **group_short_view** (bool, nullable, omitempty) — Short view in group list.

* **icons** ( [IconData](#IconData) , nullable) — Icons data.

* **is_archive** (bool, omitempty) — Contact deleted.

* **jid** ( [JID](#JID) ) — Contact Id.

* **last_activity** (string, nullable) — Last activity in this team (iso datetime).

* **munread_first** (bool, nullable, omitempty) — Show unread chats first in feed in mobile app.

* **mood** (string, omitempty) — Mood in this team.

* **patronymic** (string, nullable, omitempty) — Patronymic, if any.

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



### ContactCustomFields
Extra contact fields

* **company** (string, omitempty) — DOCUMENTATION MISSING.

* **department** (string, omitempty) — DOCUMENTATION MISSING.

* **mobile_phone** (string, omitempty) — DOCUMENTATION MISSING.

* **source** (string, omitempty) — DOCUMENTATION MISSING.

* **title** (string, omitempty) — DOCUMENTATION MISSING.



### ContactPreview
MISSING CLASS DOCUMENTATION

* **_error** (string, omitempty) — DOCUMENTATION MISSING.

* **family_name** (string) — DOCUMENTATION MISSING.

* **given_name** (string) — DOCUMENTATION MISSING.

* **patronymic** (string, omitempty) — DOCUMENTATION MISSING.

* **phone** (string) — DOCUMENTATION MISSING.

* **role** (string) — DOCUMENTATION MISSING.

* **section** (string) — DOCUMENTATION MISSING.



### ContactShort
Short contact representation

* **display_name** (string) — Full name in chats.

* **icons** ( [IconData](#IconData) , nullable) — Icons data.

* **jid** ( [JID](#JID) ) — Contact Id.

* **short_name** (string) — Short name in chats.



### Country
Country for phone numbers selection on login screen

* **code** (string) — Country code.

* **default** (bool, omitempty) — Selected by default.

* **name** (string) — Country name.

* **popular** (bool, omitempty) — Is popular, need to cache.



### DeletedChat
Minimal chat representation for deletion

* **chat_type** ( [ChatType](#ChatType) ) — Chat type.

* **gentime** (int64) — Chat fields (related to concrete participant) version.

* **is_archive** (bool) — Archive flag. Always true for this structure.

* **jid** ( [JID](#JID) ) — Group/Task/Contact id.



### DeletedRemind
Remind deleted message

* **uid** (string) — Remind id.



### DeletedSection
MISSING CLASS DOCUMENTATION

* **gentime** (int64) — Object version.

* **uid** (string) — Section uid.



### DeletedTag
Delete tag message

* **uid** (string) — Tag id.



### DeletedTeam
Team deletion message. Readonly

* **gentime** (int64) — Object version.

* **is_archive** (bool) — Team deleted.

* **uid** (string) — Team id.



### Dist
MISSING CLASS DOCUMENTATION

* **type** (string) — DOCUMENTATION MISSING.

* **url** (string) — DOCUMENTATION MISSING.



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

* **installation_type** (string) — Name of installation.

* **is_testing** (bool) — Testing installation.

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



### FontColors
MISSING CLASS DOCUMENTATION

* **brand_button** (string) — DOCUMENTATION MISSING.

* **bubble_received** (string) — DOCUMENTATION MISSING.

* **bubble_sent** (string) — DOCUMENTATION MISSING.

* **simple_button** (string) — DOCUMENTATION MISSING.

* **sub** (string) — DOCUMENTATION MISSING.

* **text** (string) — DOCUMENTATION MISSING.

* **title** (string) — DOCUMENTATION MISSING.



### GroupAccessRequest
MISSING CLASS DOCUMENTATION

* **created** (ISODateTimeString) — DOCUMENTATION MISSING.

* **subject** ( [JID](#JID) , nullable) — DOCUMENTATION MISSING.

* **uid** (string) — DOCUMENTATION MISSING.



### GroupMembership
Group chat membership status

* **can_remove** (bool, omitempty) — Can I remove this member.

* **jid** ( [JID](#JID) , nullable) — Contact id.

* **status** ( [GroupStatus](#GroupStatus) ) — Status in group.



### ICEServer
Interactive Connectivity Establishment Server for WEB Rtc connection. Readonly

* **urls** (string) — URls.



### IconColors
MISSING CLASS DOCUMENTATION

* **brand** (string) — DOCUMENTATION MISSING.

* **other** (string) — DOCUMENTATION MISSING.

* **title** (string) — DOCUMENTATION MISSING.



### IconData
Icon data. Contains sm+lg (for uploaded image) OR stub+letters+color (for icon generated from display name)

* **color** (string, omitempty) — Stub icon background color.

* **letters** (string, omitempty) — Letters from stub icon.

* **lg** ( [SingleIcon](#SingleIcon) , nullable, omitempty) — Large image.

* **sm** ( [SingleIcon](#SingleIcon) , nullable, omitempty) — Small icon.

* **stub** (string, omitempty) — Generated image with 1-2 letters.



### InputColors
MISSING CLASS DOCUMENTATION

* **active** (string) — DOCUMENTATION MISSING.

* **disable** (string) — DOCUMENTATION MISSING.

* **error** (string) — DOCUMENTATION MISSING.

* **static** (string) — DOCUMENTATION MISSING.



### Integration
Integration for concrete chat

* **comment** (string) — Comment, if any.

* **created** (ISODateTimeString, omitempty) — Creation datetime, iso.

* **enabled** (bool) — Integration enabled.

* **form** ( [IntegrationForm](#IntegrationForm) ) — Integration form.

* **group** ( [JID](#JID) ) — Chat id.

* **help** (string, omitempty) — Full description.

* **kind** (string) — Unique integration name.

* **Title** (string) — DOCUMENTATION MISSING.

* **uid** (string, omitempty) — Id.



### IntegrationField
Integration form field

* **label** (string) — Label.

* **readonly** (bool) — Is field readonly.

* **value** (string) — Current value.



### IntegrationForm
Integration form

* **api_key** ( [IntegrationField](#IntegrationField) , nullable, omitempty) — Api key field, if any.

* **url** ( [IntegrationField](#IntegrationField) , nullable, omitempty) — Url, if any.

* **webhook_url** ( [IntegrationField](#IntegrationField) , nullable, omitempty) — Webhook url, if any.



### IntegrationKind
Integration kind

* **description** (string) — Plugin description.

* **icon** (string) — Path to icon.

* **kind** (string) — Integration unique name.

* **template** ( [Integration](#Integration) ) — Integration template.

* **title** (string) — Plugin title.



### Integrations
Complete integrations data, as received from server

* **integrations** ( [Integration](#Integration) , list) — Currently existing integrations.

* **kinds** ( [IntegrationKind](#IntegrationKind) , list) — Types of integrations available for setup.



### Invitation
MISSING CLASS DOCUMENTATION

* **created** (ISODateTimeString) — DOCUMENTATION MISSING.

* **qr** (string) — DOCUMENTATION MISSING.

* **token** (string) — DOCUMENTATION MISSING.

* **uid** (string) — DOCUMENTATION MISSING.



### JID
MISSING CLASS DOCUMENTATION

* **val** (string) — DOCUMENTATION MISSING.



### JSEP
MISSING CLASS DOCUMENTATION

* **sdp** (string) — DOCUMENTATION MISSING.

* **type** (string) — DOCUMENTATION MISSING.



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



### MessageColors
MISSING CLASS DOCUMENTATION

* **allocated** (string) — DOCUMENTATION MISSING.

* **bubble_important** (string) — DOCUMENTATION MISSING.

* **bubble_received** (string) — DOCUMENTATION MISSING.

* **bubble_sent** (string) — DOCUMENTATION MISSING.

* **status_bubble** (string) — DOCUMENTATION MISSING.

* **status_feed** (string) — DOCUMENTATION MISSING.



### MessageContent
Chat message content

* **actor** ( [JID](#JID) , nullable, omitempty) — Change actor contact id (for "change" mediatype).

* **animated** (bool, omitempty) — Upload is animated image, if any. Deprecated: use Uploads instead.

* **comment** (string, omitempty) — Comment. For audio message.

* **duration** (uint, nullable, omitempty) — Upload duration, if any. Deprecated: use Uploads instead.

* **emails** (string, nullable, list, omitempty) — Emails list (for "contact" mediatype).

* **family_name** (string, nullable, omitempty) — Family name (for "contact" mediatype).

* **given_name** (string, nullable, omitempty) — Given name (for "contact" mediatype).

* **mediaURL** (string, omitempty) — Upload url, if any. Deprecated: use Uploads instead.

* **name** (string, omitempty) — Upload name, if any. Deprecated: use Uploads instead.

* **new** (string, nullable, omitempty) — Change new value (for "change" mediatype).

* **old** (string, nullable, omitempty) — Change old value (for "change" mediatype).

* **patronymic** (string, nullable, omitempty) — Patronymic name (for "contact" mediatype).

* **pdf_version** ( [PdfVersion](#PdfVersion) , nullable, omitempty) — Pdf version, if any.

* **phones** (string, nullable, list, omitempty) — Contact phones list (for "contact" mediatype).

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



### MessageLink
Checked message links. In short: "Click here: {link.Pattern}" => "Click here: <a href='{link.Url}'>{link.Text}</a>"

* **nopreview** (bool, omitempty) — Website previews disabled.

* **pattern** (string) — Text fragment that should be replaced by link.

* **preview** ( [MessageLinkPreview](#MessageLinkPreview) , nullable, omitempty) — Optional preview info, for websites.

* **text** (string) — Text replacement.

* **uploads** ( [Upload](#Upload) , list, omitempty) — Optional upload info.

* **url** (string) — Internal or external link.

* **youtube_id** (string, omitempty) — Optional youtube movie id.



### MessageLinkPreview
Website title and description

* **description** (string, omitempty) — Website description.

* **title** (string) — Website title or og:title content.



### MessagePush
Push message over websockets. Readonly

* **chat** ( [JID](#JID) , nullable) — Chat id.

* **click_action** (string) — Url opened on click.

* **created** (ISODateTimeString) — Message creation iso datetime.

* **icon_url** (string) — Absolute url to push icon.

* **message** (string) — Push body.

* **message_id** (string) — Message id.

* **sender** ( [JID](#JID) , nullable) — Sender contact id.

* **subtitle** (string) — Push subtitle.

* **tag** (string) — Push tag (for join pushes).

* **team** (string) — Team uid.

* **title** (string) — Push title.



### MessageReaction
Message emoji reaction

* **counter** (int) — Number of reactions.

* **details** ( [MessageReactionDetail](#MessageReactionDetail) , list) — Details.

* **name** (string) — Emoji.



### MessageReactionDetail
Message reaction detail

* **created** (ISODateTimeString) — When reaction added, iso datetime.

* **name** (string) — Reaction emoji.

* **sender** ( [JID](#JID) , nullable) — Reaction author.



### OAuthService
OAuth service

* **name** (string) — Integration title.

* **url** (string) — Redirect url.



### OnlineCall
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **online_count** (int, omitempty) — Number participants in call.

* **start** (ISODateTimeString, nullable, omitempty) — Call start.

* **uid** (string) — Call id.



### OnlineContact
MISSING CLASS DOCUMENTATION

* **afk** (bool, omitempty) — Is away from keyboard.

* **jid** ( [JID](#JID) ) — Contact id.

* **mobile** (bool) — Is mobile client.



### PaginatedChats
MISSING CLASS DOCUMENTATION

* **contacts** ( [Contact](#Contact) , list, omitempty) — DOCUMENTATION MISSING.

* **count** (int) — DOCUMENTATION MISSING.

* **limit** (int) — DOCUMENTATION MISSING.

* **objects** ( [Chat](#Chat) , list) — DOCUMENTATION MISSING.

* **offset** (int) — DOCUMENTATION MISSING.



### PaginatedContacts
MISSING CLASS DOCUMENTATION

* **count** (int) — DOCUMENTATION MISSING.

* **limit** (int) — DOCUMENTATION MISSING.

* **objects** ( [Contact](#Contact) , list) — DOCUMENTATION MISSING.

* **offset** (int) — DOCUMENTATION MISSING.



### PaginatedMessages
MISSING CLASS DOCUMENTATION

* **count** (int) — DOCUMENTATION MISSING.

* **limit** (int) — DOCUMENTATION MISSING.

* **objects** ( [Message](#Message) , list) — DOCUMENTATION MISSING.

* **offset** (int) — DOCUMENTATION MISSING.



### PaginatedUploadShortMessages
MISSING CLASS DOCUMENTATION

* **count** (int) — DOCUMENTATION MISSING.

* **limit** (int) — DOCUMENTATION MISSING.

* **objects** ( [UploadShortMessage](#UploadShortMessage) , list) — DOCUMENTATION MISSING.

* **offset** (int) — DOCUMENTATION MISSING.



### PdfVersion
PDF preview of mediafile. Experimental

* **text_preview** (string, omitempty) — First string of text content.

* **url** (string) — Absolute url.



### PushDevice
MISSING CLASS DOCUMENTATION

* **allowed_notifications** (bool) — DOCUMENTATION MISSING.

* **data_badges** (bool) — DOCUMENTATION MISSING.

* **data_pushes** (bool) — DOCUMENTATION MISSING.

* **device_id** (string) — DOCUMENTATION MISSING.

* **name** (string) — DOCUMENTATION MISSING.

* **notification_token** (string) — DOCUMENTATION MISSING.

* **type** (string) — DOCUMENTATION MISSING.

* **voip_notification_token** (string) — DOCUMENTATION MISSING.



### ReceivedMessage
MISSING CLASS DOCUMENTATION

* **chat** ( [JID](#JID) ) — Chat or contact id.

* **_debug** (string, omitempty) — Debug message, if any.

* **message_id** (string) — Message id.

* **num_received** (int, omitempty) — Number of contacts received this message. Experimental.

* **received** (bool) — Is received.



### Remind
Remind

* **chat** ( [JID](#JID) , nullable) — Chat id.

* **comment** (string, omitempty) — Comment, if any.

* **fire_at** (string) — Activation time, iso.

* **uid** (string) — Remind id.



### Section
MISSING CLASS DOCUMENTATION

* **description** (string, omitempty) — Description, if any.

* **gentime** (int64) — Object version.

* **is_archive** (bool, omitempty) — Is deleted.

* **name** (string) — Name.

* **sort_ordering** (uint) — Sort ordering.

* **uid** (string) — Section uid.



### ServerCallAnswer
Call parameters

* **params** ( [serverCallAnswerParams](#serverCallAnswerParams) ) — DOCUMENTATION MISSING.



### ServerCallBuzz
Call buzzing

* **params** ( [serverCallBuzzParams](#serverCallBuzzParams) ) — DOCUMENTATION MISSING.



### ServerCallBuzzcancel
Call cancelled on buzzing

* **params** ( [serverCallBuzzcancelParams](#serverCallBuzzcancelParams) ) — DOCUMENTATION MISSING.



### ServerCallCheckFingerprint
Experimental function

* **params** ( [serverCallCheckFingerprintParams](#serverCallCheckFingerprintParams) ) — DOCUMENTATION MISSING.



### ServerCallLeave
Participant leave a call

* **params** ( [serverCallLeaveParams](#serverCallLeaveParams) ) — DOCUMENTATION MISSING.



### ServerCallMuteall
All participants in call muted

* **params** ( [serverCallMuteallParams](#serverCallMuteallParams) ) — DOCUMENTATION MISSING.



### ServerCallReject
Call rejected

* **params** ( [serverCallRejectParams](#serverCallRejectParams) ) — DOCUMENTATION MISSING.



### ServerCallRestart
Call restarted

* **params** ( [serverCallRestartParams](#serverCallRestartParams) ) — DOCUMENTATION MISSING.



### ServerCallSound
Mute/unmute call participant

* **params** ( [serverCallSoundParams](#serverCallSoundParams) ) — DOCUMENTATION MISSING.



### ServerCallState
Call participant number or parameters changed

* **params** ( [serverCallStateParams](#serverCallStateParams) ) — DOCUMENTATION MISSING.



### ServerCallTalking
Someone talks in call

* **params** ( [serverCallTalkingParams](#serverCallTalkingParams) ) — DOCUMENTATION MISSING.



### ServerChatComposing
Someone typing or recording audiomessage in chat

* **params** ( [serverChatComposingParams](#serverChatComposingParams) ) — DOCUMENTATION MISSING.



### ServerChatDeleted
Chat deleted

* **params** ( [serverChatDeletedParams](#serverChatDeletedParams) ) — DOCUMENTATION MISSING.



### ServerChatDraft
Changed draft message in chan

* **params** ( [serverChatDraftParams](#serverChatDraftParams) ) — DOCUMENTATION MISSING.



### ServerChatLastread
Changed last read message in chat

* **params** ( [serverChatLastreadParams](#serverChatLastreadParams) ) — DOCUMENTATION MISSING.



### ServerChatUpdated
Chat created or updated

* **params** ( [serverChatUpdatedParams](#serverChatUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerConfirm
Server confirmed client message

* **params** ( [serverConfirmParams](#serverConfirmParams) ) — DOCUMENTATION MISSING.



### ServerContactUpdated
Contact created or updated

* **params** ( [serverContactUpdatedParams](#serverContactUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerDebug
Debug message

* **params** ( [serverDebugParams](#serverDebugParams) ) — DOCUMENTATION MISSING.



### ServerLogin
Login from other device

* **params** ( [serverLoginParams](#serverLoginParams) ) — DOCUMENTATION MISSING.



### ServerMessagePush
Push replacement for desktop application

* **params** ( [MessagePush](#MessagePush) ) — DOCUMENTATION MISSING.



### ServerMessageReceived
Message received by someone in chat

* **params** ( [serverMessageReceivedParams](#serverMessageReceivedParams) ) — DOCUMENTATION MISSING.



### ServerMessageUpdated
Chat message created, updated or deleted

* **params** ( [serverMessageUpdatedParams](#serverMessageUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerOnline
Online team members and current active calls

* **params** ( [serverOnlineParams](#serverOnlineParams) ) — DOCUMENTATION MISSING.



### ServerPanic
Critical server error

* **params** ( [serverPanicParams](#serverPanicParams) ) — DOCUMENTATION MISSING.



### ServerProcessing
Status of background operation

* **params** ( [serverProcessingParams](#serverProcessingParams) ) — DOCUMENTATION MISSING.



### ServerRemindDeleted
Task or group remind deleted

* **params** ( [serverRemindDeletedParams](#serverRemindDeletedParams) ) — DOCUMENTATION MISSING.



### ServerRemindFired
Task or group remind fired

* **params** ( [serverRemindFiredParams](#serverRemindFiredParams) ) — DOCUMENTATION MISSING.



### ServerRemindUpdated
Task/group remind created or changed

* **params** ( [serverRemindUpdatedParams](#serverRemindUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerSectionDeleted
Contact section or task project deleted

* **params** ( [serverSectionDeletedParams](#serverSectionDeletedParams) ) — DOCUMENTATION MISSING.



### ServerSectionUpdated
Contact section or task project created or changed

* **params** ( [serverSectionUpdatedParams](#serverSectionUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerTagDeleted
Tag deleted

* **params** ( [serverTagDeletedParams](#serverTagDeletedParams) ) — DOCUMENTATION MISSING.



### ServerTagUpdated
Tag created or changed

* **params** ( [serverTagUpdatedParams](#serverTagUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerTeamCounters
Counters form other teams

* **params** ( [serverTeamCountersParams](#serverTeamCountersParams) ) — DOCUMENTATION MISSING.



### ServerTeamDeleted
Team archived

* **params** ( [serverTeamDeletedParams](#serverTeamDeletedParams) ) — DOCUMENTATION MISSING.



### ServerTeamUpdated
Team created or changed

* **params** ( [serverTeamUpdatedParams](#serverTeamUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerTime
Current server time

* **params** ( [serverTimeParams](#serverTimeParams) ) — DOCUMENTATION MISSING.



### ServerUiSettings
Part of UI settings changed

* **params** ( [UiSettings](#UiSettings) , nullable) — DOCUMENTATION MISSING.



### ServerUploadUpdated
Upload object created or changed

* **params** ( [serverUploadUpdatedParams](#serverUploadUpdatedParams) ) — DOCUMENTATION MISSING.



### ServerWarning
Something went wrong with client message

* **params** ( [serverWarningParams](#serverWarningParams) ) — DOCUMENTATION MISSING.



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



### SingleIcon
Small or large icon

* **height** (int) — Icon height, in pixels.

* **url** (string) — absolute url to icon.

* **width** (int) — Icon width, in pixels.



### Sticker
MISSING CLASS DOCUMENTATION

* **icon100** (string) — DOCUMENTATION MISSING.

* **icon128** (string) — DOCUMENTATION MISSING.

* **icon200** (string) — DOCUMENTATION MISSING.

* **icon64** (string) — DOCUMENTATION MISSING.

* **message_content** ( [MessageContent](#MessageContent) ) — DOCUMENTATION MISSING.

* **uid** (string) — DOCUMENTATION MISSING.



### Stickerpack
MISSING CLASS DOCUMENTATION

* **author** (string, omitempty) — DOCUMENTATION MISSING.

* **author_link** (string, omitempty) — DOCUMENTATION MISSING.

* **name** (string) — DOCUMENTATION MISSING.

* **stickers** ( [Sticker](#Sticker) , list) — DOCUMENTATION MISSING.

* **title** (string) — DOCUMENTATION MISSING.

* **uid** (string) — DOCUMENTATION MISSING.



### Subtask
Link to sub/sup task

* **assignee** ( [JID](#JID) ) — Assignee contact id. Tasks only.

* **display_name** (string) — Title.

* **jid** ( [JID](#JID) ) — Task id.

* **num** (uint) — Task number in this team.

* **public** (bool, omitempty) — Can other team member see this task/group chat.

* **title** (string) — Task title. Generated from number and description.



### SwitcherColors
MISSING CLASS DOCUMENTATION

* **off** (string) — DOCUMENTATION MISSING.

* **on** (string) — DOCUMENTATION MISSING.



### Tag
Task tag

* **name** (string) — Tag name.

* **uid** (string) — Tag id.



### TaskColor
Task color rules color

* **dark** (string) — DOCUMENTATION MISSING.

* **light** (string) — DOCUMENTATION MISSING.

* **regular** (string) — DOCUMENTATION MISSING.



### TaskCounters
Tasks counters

* **jid** ( [JID](#JID) ) — Task jid.

* **num_unread** (uint, omitempty) — Unreads counter.

* **num_unread_notices** (uint, omitempty) — Mentions (@) counter.



### TaskFilter
Task filter

* **field** ( [TaskFilterKey](#TaskFilterKey) ) — Task filter field.

* **title** (string) — Filter title.



### TaskItem
Task checklist item

* **can_toggle** (bool, omitempty) — Can I toggle this item.

* **checked** (bool, omitempty) — Item checked.

* **sort_ordering** (uint, omitempty) — Sort ordering.

* **subtask** ( [Subtask](#Subtask) , nullable, omitempty) — Link to subtask. Optional.

* **text** (string) — Text or "#{OtherTaskNumber}".

* **uid** (string, omitempty) — Id.



### TaskItems
MISSING CLASS DOCUMENTATION

* **checked** (bool) — DOCUMENTATION MISSING.

* **name** (string) — DOCUMENTATION MISSING.



### TaskPreview
MISSING CLASS DOCUMENTATION

* **assignee** ( [JID](#JID) ) — DOCUMENTATION MISSING.

* **deadline** (time.Time, nullable) — DOCUMENTATION MISSING.

* **description** (string) — DOCUMENTATION MISSING.

* **_error** (string, omitempty) — DOCUMENTATION MISSING.

* **items** ( [TaskItems](#TaskItems) , list) — DOCUMENTATION MISSING.

* **public** (bool) — DOCUMENTATION MISSING.

* **section** (string) — DOCUMENTATION MISSING.

* **tags** (string, list) — DOCUMENTATION MISSING.



### TaskSort
Task sort type

* **key** ( [TaskSortKey](#TaskSortKey) ) — Field.

* **title** (string) — Sort title.



### TaskStatus
Custom task status

* **is_archive** (bool, omitempty) — Status not used anymore.

* **name** (string) — Status internal name.

* **sort_ordering** (uint) — Status sort ordering.

* **title** (string) — Status localized name.

* **uid** (string, omitempty) — Status id.



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

* **single_group** ( [JID](#JID) , readonly for clients, nullable, omitempty) — For single group teams, jid of chat.

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



### TeamCounter
Unread message counters

* **uid** (string) — Team id.

* **unread** ( [TeamUnread](#TeamUnread) ) — Unread message counters.



### TeamShort
Short team representation. For invites, push notifications, etc. Readonly

* **icons** ( [IconData](#IconData) ) — Team icons.

* **name** (string) — Team name.

* **uid** (string) — Team id.



### Terms
Experimental translation fields for "team" entity renaming. Readonly

* **EnInTeam** (string) — DOCUMENTATION MISSING.

* **EnTeam** (string) — DOCUMENTATION MISSING.

* **EnTeamAccess** (string) — DOCUMENTATION MISSING.

* **EnTeamAdmin** (string) — DOCUMENTATION MISSING.

* **EnTeamAdmins** (string) — DOCUMENTATION MISSING.

* **EnTeamGuest** (string) — DOCUMENTATION MISSING.

* **EnTeamMember** (string) — DOCUMENTATION MISSING.

* **EnTeamMembers** (string) — DOCUMENTATION MISSING.

* **EnTeamOwner** (string) — DOCUMENTATION MISSING.

* **EnTeamSettings** (string) — DOCUMENTATION MISSING.

* **EnTeams** (string) — DOCUMENTATION MISSING.

* **EnToTeam** (string) — DOCUMENTATION MISSING.

* **RuInTeam** (string) — DOCUMENTATION MISSING.

* **RuTeam** (string) — DOCUMENTATION MISSING.

* **RuTeamAccess** (string) — DOCUMENTATION MISSING.

* **RuTeamAdmin** (string) — DOCUMENTATION MISSING.

* **RuTeamAdmins** (string) — DOCUMENTATION MISSING.

* **RuTeamD** (string) — DOCUMENTATION MISSING.

* **RuTeamGuest** (string) — DOCUMENTATION MISSING.

* **RuTeamMember** (string) — DOCUMENTATION MISSING.

* **RuTeamMembers** (string) — DOCUMENTATION MISSING.

* **RuTeamOwner** (string) — DOCUMENTATION MISSING.

* **RuTeamP** (string) — DOCUMENTATION MISSING.

* **RuTeamR** (string) — DOCUMENTATION MISSING.

* **RuTeamSettings** (string) — DOCUMENTATION MISSING.

* **RuTeamT** (string) — DOCUMENTATION MISSING.

* **RuTeamV** (string) — DOCUMENTATION MISSING.

* **RuTeams** (string) — DOCUMENTATION MISSING.

* **RuTeamsD** (string) — DOCUMENTATION MISSING.

* **RuTeamsP** (string) — DOCUMENTATION MISSING.

* **RuTeamsR** (string) — DOCUMENTATION MISSING.

* **RuTeamsT** (string) — DOCUMENTATION MISSING.

* **RuTeamsV** (string) — DOCUMENTATION MISSING.

* **RuToTeam** (string) — DOCUMENTATION MISSING.



### Theme
Color theme

* **AccentColor** (string) — DOCUMENTATION MISSING.

* **AccentHoverColor** (string) — DOCUMENTATION MISSING.

* **AppAccentColor** (string) — Deprecated.

* **AppPrimaryColor** (string) — Deprecated.

* **attention** (string) — DOCUMENTATION MISSING.

* **attention_light** (string) — DOCUMENTATION MISSING.

* **back** (string) — DOCUMENTATION MISSING.

* **back_dark** (string) — DOCUMENTATION MISSING.

* **back_light** (string) — DOCUMENTATION MISSING.

* **background** (string) — DOCUMENTATION MISSING.

* **BgColor** (string) — Web colors.

* **BgHoverColor** (string) — DOCUMENTATION MISSING.

* **brand** (string) — App colors.

* **brand_dark** (string) — DOCUMENTATION MISSING.

* **brand_light** (string) — DOCUMENTATION MISSING.

* **button** ( [ButtonColors](#ButtonColors) , nullable) — DOCUMENTATION MISSING.

* **chat_input_background** (string) — DOCUMENTATION MISSING.

* **error** (string) — DOCUMENTATION MISSING.

* **error_light** (string) — DOCUMENTATION MISSING.

* **font** ( [FontColors](#FontColors) , nullable) — DOCUMENTATION MISSING.

* **ic** ( [IconColors](#IconColors) , nullable) — DOCUMENTATION MISSING.

* **input** ( [InputColors](#InputColors) , nullable) — DOCUMENTATION MISSING.

* **MainAccent** (string) — DOCUMENTATION MISSING.

* **MainAccentHover** (string) — DOCUMENTATION MISSING.

* **MainLightAccent** (string) — DOCUMENTATION MISSING.

* **MainLink** (string) — DOCUMENTATION MISSING.

* **message** ( [MessageColors](#MessageColors) , nullable) — DOCUMENTATION MISSING.

* **modal_background** (string) — DOCUMENTATION MISSING.

* **MutedTextColor** (string) — DOCUMENTATION MISSING.

* **substrate_background** (string) — DOCUMENTATION MISSING.

* **success** (string) — DOCUMENTATION MISSING.

* **success_light** (string) — DOCUMENTATION MISSING.

* **switcher** ( [SwitcherColors](#SwitcherColors) , nullable) — DOCUMENTATION MISSING.

* **tab_background** (string) — DOCUMENTATION MISSING.

* **TextColor** (string) — DOCUMENTATION MISSING.

* **TextOnAccentHoverColor** (string) — DOCUMENTATION MISSING.

* **title_background** (string) — DOCUMENTATION MISSING.



### Unread
Unread message counters

* **chats** (uint) — Total chats with unread messages.

* **messages** (uint) — Total unread messages.

* **notice_messages** (uint) — Total unread messages with mentions.



### Upload
Uploaded media

* **animated** (bool, omitempty) — Is animated (images only).

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



### UploadPreview
Upload preview

* **height** (int) — Height in pixels.

* **url** (string) — Absolute url to image.

* **url_2x** (string) — Absolute url to high resolution image (retina).

* **width** (int) — Width in pixels.



### UploadShortMessage
Upload + ShortMessage

* **message** ( [ShortMessage](#ShortMessage) ) — DOCUMENTATION MISSING.

* **upload** ( [Upload](#Upload) ) — DOCUMENTATION MISSING.



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



### UserAuth
MISSING CLASS DOCUMENTATION

* **addr** (string) — DOCUMENTATION MISSING.

* **created** (time.Time) — DOCUMENTATION MISSING.

* **_age** (int, omitempty) — DOCUMENTATION MISSING.

* **device** ( [PushDevice](#PushDevice) , nullable, omitempty) — DOCUMENTATION MISSING.

* **kind** (string) — DOCUMENTATION MISSING.

* **last_access** (time.Time, omitempty) — DOCUMENTATION MISSING.

* **uid** (string) — DOCUMENTATION MISSING.

* **user_agent** (string) — DOCUMENTATION MISSING.



### UserWithMe
MISSING CLASS DOCUMENTATION

* **devices** ( [PushDevice](#PushDevice) , list) — DOCUMENTATION MISSING.

* **inviter** ( [JID](#JID) , nullable, omitempty) — DOCUMENTATION MISSING.

* **teams** ( [Team](#Team) , list) — DOCUMENTATION MISSING.



### Wallpaper
Chat wallpaper

* **key** (string) — Unique identifier.

* **name** (string) — Localized description.

* **url** (string) — Url to jpg or png.



### WikiPage
Wiki page. Experimental

* **editor** ( [JID](#JID) ) — Last editor contact id.

* **gentime** (int64) — Object version.

* **text** (string) — Page text.

* **updated** (string) — Update time, iso.



### clientActivityParams
MISSING CLASS DOCUMENTATION

* **afk** (bool) — Is away from keyboard.



### clientCallBuzzCancelParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.



### clientCallBuzzParams
Call buzzing

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **members** ( [JID](#JID) , list) — List of call participants.



### clientCallLeaveParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **reason** (string) — Reason, if any.



### clientCallMuteAllParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.



### clientCallOfferParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **muted** (bool) — Mute state.

* **sdp** (string) — SDP (session description protocol) data.

* **trickle** (bool) — Is trickle mode enabled.



### clientCallRejectParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **reason** (string) — Reason, if any.



### clientCallSoundParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **muted** (bool) — Mute state.



### clientCallTrickleParams
MISSING CLASS DOCUMENTATION

* **candidate** (string) — Trickle candidate.

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **sdp_mid** (string) — SDP mid.

* **sdp_mline_index** (int) — SDP index.



### clientChatComposingParams
MISSING CLASS DOCUMENTATION

* **composing** (bool, omitempty) — true = start typing / audio recording, false = stop.

* **draft** (string, nullable, omitempty) — Message draft data.

* **is_audio** (bool, omitempty) — true = audiomessage, false = text typing.

* **jid** ( [JID](#JID) ) — Chat or contact id.



### clientChatLastreadParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **last_read_message_id** (string, nullable, omitempty) — Last read message id. Omitted = last message in chat.



### clientConfirmParams
MISSING CLASS DOCUMENTATION

* **confirm_id** (string) — Unique identifier generated by client.



### clientMessageDeletedParams
MISSING CLASS DOCUMENTATION

* **message_id** (string, omitempty) — Message id.



### serverCallAnswerCandidate
MISSING CLASS DOCUMENTATION

* **candidate** (string) — DOCUMENTATION MISSING.

* **sdpMLineIndex** (int) — DOCUMENTATION MISSING.



### serverCallAnswerParams
MISSING CLASS DOCUMENTATION

* **candidates** ( [serverCallAnswerCandidate](#serverCallAnswerCandidate) , list) — List of ICE candidates (when trickle = false).

* **jsep** ( [JSEP](#JSEP) ) — SDP data.

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **onliners** ( [CallOnliner](#CallOnliner) , list) — Current call participants.

* **uid** (string) — Call id.



### serverCallBuzzParams
MISSING CLASS DOCUMENTATION

* **actor** ( [ContactShort](#ContactShort) ) — Short call creator information.

* **buzz_timeout** (int) — Number of seconds for stop buzzing.

* **chat** ( [ChatShort](#ChatShort) ) — Short chat information.

* **display_name** (string) — Chat title.

* **icons** ( [IconData](#IconData) , nullable) — Chat icons.

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **team** (string) — Deprecated.

* **teaminfo** ( [TeamShort](#TeamShort) ) — Short team information.

* **uid** (string) — Call id.



### serverCallBuzzcancelParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **team** (string) — Team id.

* **uid** (string) — Call id.



### serverCallCheckFingerprintParams
MISSING CLASS DOCUMENTATION

* **fingerprint** (string) — DOCUMENTATION MISSING.



### serverCallLeaveParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **uid** (string) — Call uid.



### serverCallMuteallParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **muted** (bool) — Mute state.



### serverCallRejectParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **reason** (string) — Reason, if any.

* **uid** (string) — Call id.



### serverCallRestartParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **team** (string) — Team id.

* **uid** (string) — Call id.



### serverCallSoundParams
MISSING CLASS DOCUMENTATION

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **muted** (bool) — Mute state.



### serverCallStateParams
MISSING CLASS DOCUMENTATION

* **audiorecord** (bool) — Call record enabled.

* **buzz** (bool, omitempty) — Call buzzing.

* **finish** (ISODateTimeString, nullable) — Call finish, if any.

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **onliners** ( [CallOnliner](#CallOnliner) , list) — Call participants.

* **start** (ISODateTimeString, nullable) — Call start, if any.

* **timestamp** (int64) — Event start. FIXME: why not gentime?.

* **uid** (string) — Call id.



### serverCallTalkingParams
MISSING CLASS DOCUMENTATION

* **actor** ( [JID](#JID) ) — Actor id.

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **talking** (bool) — Is talking.



### serverChatComposingParams
MISSING CLASS DOCUMENTATION

* **actor** ( [JID](#JID) ) — Actor id.

* **composing** (bool) — true = start typing / audio recording, false = stop.

* **is_audio** (bool, omitempty) — true = audiomessage, false = text typing.

* **jid** ( [JID](#JID) ) — Chat or contact id.

* **valid_until** (ISODateTimeString, omitempty) — Composing event max lifetime.



### serverChatDeletedParams
MISSING CLASS DOCUMENTATION

* **badge** (uint) — Total number of unreads.

* **chats** ( [DeletedChat](#DeletedChat) , list) — List of deleted chats.

* **team_unread** ( [TeamUnread](#TeamUnread) , nullable) — Current team counters.



### serverChatDraftParams
MISSING CLASS DOCUMENTATION

* **draft** (string) — Draft text.

* **draft_num** (int64) — Draft version.

* **jid** ( [JID](#JID) , nullable) — Chat or contact id.



### serverChatLastreadParams
MISSING CLASS DOCUMENTATION

* **badge** (uint) — Total number of unreads.

* **chats** ( [ChatCounters](#ChatCounters) , list) — Chat counters.

* **team_unread** ( [TeamUnread](#TeamUnread) , nullable) — Current team counters.



### serverChatUpdatedParams
MISSING CLASS DOCUMENTATION

* **badge** (uint) — Total number of unreads.

* **chats** ( [Chat](#Chat) , list) — Chat counters.

* **team_unread** ( [TeamUnread](#TeamUnread) , nullable) — Current team counters.



### serverConfirmParams
MISSING CLASS DOCUMENTATION

* **confirm_id** (string) — Unique id generated by server.



### serverContactUpdatedParams
MISSING CLASS DOCUMENTATION

* **contacts** ( [Contact](#Contact) , list) — Contact info.



### serverDebugParams
MISSING CLASS DOCUMENTATION

* **text** (string) — Debug message.



### serverLoginParams
MISSING CLASS DOCUMENTATION

* **device_name** (string) — Device name.



### serverMessageReceivedParams
MISSING CLASS DOCUMENTATION

* **messages** ( [ReceivedMessage](#ReceivedMessage) , list) — received message data.



### serverMessageUpdatedParams
MISSING CLASS DOCUMENTATION

* **badge** (uint, nullable) — Total number of unreads, if changed.

* **chat_counters** ( [ChatCounters](#ChatCounters) , list) — Chat counters.

* **delayed** (bool) — true = silently message update, false = new message.

* **messages** ( [Message](#Message) , list) — Messages data.

* **team_unread** ( [TeamUnread](#TeamUnread) , nullable) — Current team counters.



### serverOnlineParams
MISSING CLASS DOCUMENTATION

* **calls** ( [OnlineCall](#OnlineCall) , nullable, list) — Active calls.

* **contacts** ( [OnlineContact](#OnlineContact) , nullable, list) — Online team members.



### serverPanicParams
MISSING CLASS DOCUMENTATION

* **code** (string) — Error code.

* **debug** (string, omitempty) — Debug message.



### serverProcessingParams
MISSING CLASS DOCUMENTATION

* **action** (string) — Action name.

* **has_error** (bool) — Has error.

* **message** (string) — Message.

* **num** (int) — Current processing item.

* **total** (int) — Total processing items.



### serverRemindDeletedParams
MISSING CLASS DOCUMENTATION

* **reminds** ( [DeletedRemind](#DeletedRemind) , list) — Remind information.



### serverRemindFiredParams
MISSING CLASS DOCUMENTATION

* **reminds** ( [Remind](#Remind) , list) — Remind information.



### serverRemindUpdatedParams
MISSING CLASS DOCUMENTATION

* **reminds** ( [Remind](#Remind) , list) — Remind information.



### serverSectionDeletedParams
MISSING CLASS DOCUMENTATION

* **chat_type** ( [ChatType](#ChatType) ) — Chat type.

* **gentime** (int64) — Deprecated.

* **sections** ( [DeletedSection](#DeletedSection) , list) — Section/project info.



### serverSectionUpdatedParams
MISSING CLASS DOCUMENTATION

* **chat_type** ( [ChatType](#ChatType) ) — Chat type.

* **gentime** (int64) — deprecated.

* **sections** ( [Section](#Section) , list) — Section/project info.



### serverTagDeletedParams
MISSING CLASS DOCUMENTATION

* **tags** ( [DeletedTag](#DeletedTag) , list) — Tags info.



### serverTagUpdatedParams
MISSING CLASS DOCUMENTATION

* **tags** ( [Tag](#Tag) , list) — Tags info.



### serverTeamCountersParams
MISSING CLASS DOCUMENTATION

* **badge** (uint) — Total number of unreads.

* **teams** ( [TeamCounter](#TeamCounter) , list) — Counters.



### serverTeamDeletedParams
MISSING CLASS DOCUMENTATION

* **teams** ( [DeletedTeam](#DeletedTeam) , list) — Teams info.



### serverTeamUpdatedParams
MISSING CLASS DOCUMENTATION

* **teams** ( [Team](#Team) , list) — DOCUMENTATION MISSING.



### serverTimeParams
MISSING CLASS DOCUMENTATION

* **time** (ISODateTimeString) — Current time.



### serverUploadUpdatedParams
MISSING CLASS DOCUMENTATION

* **uploads** ( [Upload](#Upload) , list) — Uploads data.



### serverWarningParams
MISSING CLASS DOCUMENTATION

* **message** (string) — Message.

* **orig** (interface{}) — Debug information.


