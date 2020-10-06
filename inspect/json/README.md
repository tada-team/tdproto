## Protocol description in json:
```json
[
    {
        "name": "CallEvent",
        "help": "Audiocall information",
        "fields": [
            {
                "name": "Start",
                "help": "Call start, iso date",
                "json": "start",
                "type": "string",
                "null": true
            },
            {
                "name": "Finish",
                "help": "Call finish, iso date",
                "json": "finish",
                "type": "string",
                "null": true
            },
            {
                "name": "Audiorecord",
                "help": "Call record enabled",
                "json": "audiorecord",
                "type": "bool"
            },
            {
                "name": "Onliners",
                "help": "Call members",
                "json": "onliners",
                "type": "[]CallOnliner"
            }
        ]
    },
    {
        "name": "CallOnliner",
        "help": "Call participant",
        "fields": [
            {
                "name": "Jid",
                "help": "Contact id",
                "json": "jid",
                "type": "JID"
            },
            {
                "name": "DisplayName",
                "help": "Contact name",
                "json": "display_name",
                "type": "string"
            },
            {
                "name": "Icon",
                "help": "Contact icon",
                "json": "icon",
                "type": "string"
            },
            {
                "name": "Muted",
                "help": "Microphone muted. Computed from devices muted states",
                "json": "muted",
                "type": "bool"
            },
            {
                "name": "Devices",
                "help": "Member devices, strictly one for now",
                "json": "devices",
                "type": "[]CallDevice"
            }
        ]
    },
    {
        "name": "CallDevice",
        "help": "Call participant device",
        "fields": [
            {
                "name": "Muted",
                "help": "Device muted",
                "json": "muted",
                "type": "bool"
            },
            {
                "name": "Useragent",
                "help": "Device description",
                "json": "useragent",
                "type": "string"
            }
        ]
    },
    {
        "name": "SingleIcon",
        "help": "Small or large icon",
        "fields": [
            {
                "name": "Url",
                "help": "absolute url to icon",
                "json": "url",
                "type": "string"
            },
            {
                "name": "Width",
                "help": "Icon width, in pixels",
                "json": "width",
                "type": "int"
            },
            {
                "name": "Height",
                "help": "Icon height, in pixels",
                "json": "height",
                "type": "int"
            }
        ]
    },
    {
        "name": "IconData",
        "help": "Icon data. Contains sm+lg (for uploaded image) OR stub+letters+color (for icon generated from display name)",
        "fields": [
            {
                "name": "Sm",
                "help": "Small icon",
                "json": "sm",
                "type": "SingleIcon",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Lg",
                "help": "Large image",
                "json": "lg",
                "type": "SingleIcon",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Stub",
                "help": "Generated image with 1-2 letters",
                "json": "stub",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Letters",
                "help": "Letters from stub icon",
                "json": "letters",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Color",
                "help": "Stub icon background color",
                "json": "color",
                "type": "string",
                "omitempty": true
            }
        ]
    },
    {
        "name": "ChatShort",
        "help": "Mimimal chat representaion",
        "fields": [
            {
                "name": "Jid",
                "help": "Group/Task/Contact id",
                "json": "jid",
                "type": "JID"
            },
            {
                "name": "ChatType",
                "help": "Chat type",
                "json": "chat_type",
                "type": "ChatType"
            },
            {
                "name": "DisplayName",
                "help": "Title",
                "json": "display_name",
                "type": "string"
            },
            {
                "name": "Icons",
                "help": "Icon data",
                "json": "icons",
                "type": "IconData",
                "null": true
            }
        ]
    },
    {
        "name": "DeletedChat",
        "help": "Mimimal chat representaion for deletion",
        "fields": [
            {
                "name": "Jid",
                "help": "Group/Task/Contact id",
                "json": "jid",
                "type": "JID"
            },
            {
                "name": "ChatType",
                "help": "Chat type",
                "json": "chat_type",
                "type": "ChatType"
            },
            {
                "name": "Gentime",
                "help": "Chat fields (related to concrete participan) changes indicator",
                "json": "gentime",
                "type": "int64"
            },
            {
                "name": "IsArchive",
                "help": "Archive flag. Always true for this structure",
                "json": "is_archive",
                "type": "bool"
            }
        ]
    },
    {
        "name": "Chat",
        "help": "Chat (direct, group, task) representaion",
        "fields": [
            {
                "name": "Jid",
                "help": "Group/Task/Contact id",
                "json": "jid",
                "type": "JID"
            },
            {
                "name": "ChatType",
                "help": "Chat type",
                "json": "chat_type",
                "type": "ChatType"
            },
            {
                "name": "BaseGentime",
                "help": "Base fields (not related to concrete participant) changes indicator",
                "json": "base_gentime",
                "type": "int64",
                "omitempty": true
            },
            {
                "name": "Gentime",
                "help": "Chat fields (related to concrete participan) changes indicator",
                "json": "gentime",
                "type": "int64"
            },
            {
                "name": "Created",
                "help": "Creation date, iso datetime",
                "json": "created",
                "type": "string"
            },
            {
                "name": "DisplayName",
                "help": "Title",
                "json": "display_name",
                "type": "string"
            },
            {
                "name": "Icons",
                "help": "Icons info",
                "json": "icons",
                "type": "IconData",
                "null": true
            },
            {
                "name": "CountersEnabled",
                "help": "Include unread messages to counters",
                "json": "counters_enabled",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanCall",
                "help": "Can I call to this chat",
                "json": "can_call",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanSendMessage",
                "help": "Can I send message to this chat",
                "json": "can_send_message",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CantSendMessageReason",
                "help": "Why I can't send message to this chat (if can't)",
                "json": "cant_send_message_reason",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Collapsed",
                "help": "Description collapsed. Used for tasks only",
                "json": "collapsed",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "Draft",
                "help": "Last message draft, if any",
                "json": "draft",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "DraftNum",
                "help": "Last message draft version , if any",
                "json": "draft_num",
                "type": "int64",
                "omitempty": true
            },
            {
                "name": "Hidden",
                "help": "Hidden chat",
                "json": "hidden",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "NotificationsEnabled",
                "help": "Push notifications enabled",
                "json": "notifications_enabled",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "NumImportants",
                "help": "Number of importants messages",
                "json": "num_importants",
                "type": "int",
                "omitempty": true
            },
            {
                "name": "NumUnread",
                "help": "Unreads conuter",
                "json": "num_unread",
                "type": "uint",
                "omitempty": true
            },
            {
                "name": "NumUnreadNotices",
                "help": "Mentions (@) counter",
                "json": "num_unread_notices",
                "type": "uint",
                "omitempty": true
            },
            {
                "name": "LastMessage",
                "help": "Last message object",
                "json": "last_message",
                "type": "Message",
                "null": true,
                "omitempty": true
            },
            {
                "name": "LastReadMessageId",
                "help": "Last read message id, if any",
                "json": "last_read_message_id",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Section",
                "help": "Project / section id, if any",
                "json": "section",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "ChangeableFields",
                "help": "List of editable fields",
                "json": "changeable_fields",
                "type": "[]string",
                "omitempty": true
            },
            {
                "name": "Pinned",
                "help": "Is chat pinned on top",
                "json": "pinned",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "PinnedSortOrdering",
                "help": "Sort oreding for pinned chat",
                "json": "pinned_sort_ordering",
                "type": "int",
                "omitempty": true
            },
            {
                "name": "NumMembers",
                "help": "Non-archive participants number",
                "json": "num_members",
                "type": "uint",
                "null": true,
                "omitempty": true
            },
            {
                "name": "CanDelete",
                "help": "Can I delete this chat",
                "json": "can_delete",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "Description",
                "help": "Group or task description",
                "json": "description",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Feed",
                "help": "Present in feed (main screen)",
                "json": "feed",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "PinnedMessage",
                "help": "Pinned message for this chat",
                "json": "pinned_message",
                "type": "Message",
                "null": true,
                "omitempty": true
            },
            {
                "name": "ColorIndex",
                "help": "Custom color index from table of colors. Tasks only",
                "json": "color_index",
                "type": "uint16",
                "null": true,
                "omitempty": true
            },
            {
                "name": "NumItems",
                "help": "Items in checklist. Tasks only",
                "json": "num_items",
                "type": "uint",
                "null": true,
                "omitempty": true
            },
            {
                "name": "NumCheckedItems",
                "help": "Checked items in checklist. Tasks only",
                "json": "num_checked_items",
                "type": "uint",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Assignee",
                "help": "Assignee contact id. Tasks only",
                "json": "assignee",
                "type": "JID",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Num",
                "help": "Task number in this team",
                "json": "num",
                "type": "uint",
                "omitempty": true
            },
            {
                "name": "Observers",
                "help": "Task observers id's",
                "json": "observers",
                "type": "[]JID",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Owner",
                "help": "Task creator",
                "json": "owner",
                "type": "JID",
                "null": true,
                "omitempty": true
            },
            {
                "name": "TaskStatus",
                "help": "Task status. May be custom",
                "json": "task_status",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Title",
                "help": "Task title. Generated from number and description",
                "json": "title",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Done",
                "help": "Task done date in iso format, if any",
                "json": "done",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "DoneReason",
                "help": "Task done reason, if any",
                "json": "done_reason",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "Deadline",
                "help": "Task deadline in iso format, if any",
                "json": "deadline",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "DeadlineExpired",
                "help": "Is task deadline expired",
                "json": "deadline_expired",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "Links",
                "help": "Links in description",
                "json": "links",
                "type": "MessageLinks",
                "omitempty": true
            },
            {
                "name": "Tags",
                "help": "Task tags list, if any",
                "json": "tags",
                "type": "[]string",
                "omitempty": true
            },
            {
                "name": "Importance",
                "help": "Task importance, if available in team",
                "json": "importance",
                "type": "int",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Urgency",
                "help": "Task urgency, if available in team",
                "json": "urgency",
                "type": "int",
                "null": true,
                "omitempty": true
            },
            {
                "name": "SpentTime",
                "help": "Task spent time, number",
                "json": "spent_time",
                "type": "int",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Complexity",
                "help": "Task complexity, number",
                "json": "complexity",
                "type": "int",
                "null": true,
                "omitempty": true
            },
            {
                "name": "LinkedMessages",
                "help": "Used for \"Create task from messages...\"",
                "json": "linked_messages",
                "type": "[]interface{}",
                "omitempty": true
            },
            {
                "name": "Items",
                "help": "Checklist items. Task only",
                "json": "items",
                "type": "[]TaskItem",
                "omitempty": true
            },
            {
                "name": "Parents",
                "help": "Parent tasks",
                "json": "parents",
                "type": "[]Subtask",
                "omitempty": true
            },
            {
                "name": "Tabs",
                "help": "Tab names",
                "json": "tabs",
                "type": "[]TaskTabKey",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Status",
                "help": "My status in group chat",
                "json": "status",
                "type": "GroupStatus",
                "omitempty": true
            },
            {
                "name": "Members",
                "help": "Group chat members",
                "json": "members",
                "type": "[]GroupMembership",
                "omitempty": true
            },
            {
                "name": "CanAddMember",
                "help": "Can I add member to this group chat",
                "json": "can_add_member",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanRemoveMember",
                "help": "Can I remove member from this group chat",
                "json": "can_remove_member",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanChangeMemberStatus",
                "help": "Can I change member status in this group chat",
                "json": "can_change_member_status",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanChangeSettings",
                "help": "deprecated: use changeable fields",
                "json": "can_change_settings",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "DefaultForAll",
                "help": "Any new team member will be added to this group chat",
                "json": "default_for_all",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "ReadonlyForMembers",
                "help": "Readonly for non-admins group chat (Like Channels in Telegram bug switchable)",
                "json": "readonly_for_members",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "AutocleanupAge",
                "help": "Delete messages in this chat in seconds. Experemental function",
                "json": "autocleanup_age",
                "type": "int",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Public",
                "help": "Can other team member see this task/group chat",
                "json": "public",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanJoin",
                "help": "Can I join to this public group/task",
                "json": "can_join",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanDeleteAnyMessage",
                "help": "Can I delete any message in this chat",
                "json": "can_delete_any_message",
                "type": "bool",
                "null": true,
                "omitempty": true
            },
            {
                "name": "CanSetImportantAnyMessage",
                "help": "Can I change Important flag in any message in this chat",
                "json": "can_set_important_any_message",
                "type": "bool",
                "null": true,
                "omitempty": true
            }
        ]
    },
    {
        "name": "Subtask",
        "help": "Link to sub/sup task",
        "fields": [
            {
                "name": "Jid",
                "help": "Task id",
                "json": "jid",
                "type": "JID"
            },
            {
                "name": "Assignee",
                "help": "Assignee contact id. Tasks only",
                "json": "assignee",
                "type": "JID"
            },
            {
                "name": "Title",
                "help": "Task title. Generated from number and description",
                "json": "title",
                "type": "string"
            },
            {
                "name": "Num",
                "help": "Task number in this team",
                "json": "num",
                "type": "uint"
            },
            {
                "name": "DisplayName",
                "help": "Title",
                "json": "display_name",
                "type": "string"
            },
            {
                "name": "Public",
                "help": "Can other team member see this task/group chat",
                "json": "public",
                "type": "bool",
                "omitempty": true
            }
        ]
    },
    {
        "name": "TaskItem",
        "help": "Task checklist item",
        "fields": [
            {
                "name": "Uid",
                "help": "Id",
                "json": "uid",
                "type": "string",
                "omitempty": true
            },
            {
                "name": "SortOrdering",
                "help": "Sort ordering",
                "json": "sort_ordering",
                "type": "uint",
                "omitempty": true
            },
            {
                "name": "Text",
                "help": "Text or \"#{OtherTaskNumber}\"",
                "json": "text",
                "type": "string"
            },
            {
                "name": "Checked",
                "help": "Item checked",
                "json": "checked",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "CanToggle",
                "help": "Can I toggle this item",
                "json": "can_toggle",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "Subtask",
                "help": "Link to subtask. Optional",
                "json": "subtask",
                "type": "Subtask",
                "null": true,
                "omitempty": true
            }
        ]
    },
    {
        "name": "GroupMembership",
        "help": "Group chat membership status",
        "fields": [
            {
                "name": "Jid",
                "help": "Contact id",
                "json": "jid",
                "type": "JID",
                "null": true
            },
            {
                "name": "Status",
                "help": "Status in group",
                "json": "status",
                "type": "GroupStatus"
            },
            {
                "name": "CanRemove",
                "help": "Can I remove this member",
                "json": "can_remove",
                "type": "bool",
                "omitempty": true
            }
        ]
    },
    {
        "name": "Features",
        "help": "{hostname}/features.js / {hostname}/features.json structure",
        "fields": [
            {
                "name": "Host",
                "help": "current host",
                "json": "host",
                "type": "string"
            },
            {
                "name": "Build",
                "help": "build/revision of server side",
                "json": "build",
                "type": "string"
            },
            {
                "name": "DesktopVersion",
                "help": "desktop application version",
                "json": "desktop_version",
                "type": "string"
            },
            {
                "name": "FrontVersion",
                "help": "webclient version",
                "json": "front_version",
                "type": "string"
            },
            {
                "name": "AppTitle",
                "help": "Application title",
                "json": "app_title",
                "type": "string"
            },
            {
                "name": "Userver",
                "help": "Static files server address",
                "json": "userver",
                "type": "string"
            },
            {
                "name": "IOSApp",
                "help": "Link to AppStore",
                "json": "ios_app",
                "type": "string"
            },
            {
                "name": "AndroidApp",
                "help": "Link to Google Play",
                "json": "android_app",
                "type": "string"
            },
            {
                "name": "Theme",
                "help": "Default UI theme",
                "json": "theme",
                "type": "string"
            },
            {
                "name": "MinAppVersion",
                "help": "Minimal application version required for this server. Used for breaking changes",
                "json": "min_app_version",
                "type": "string"
            },
            {
                "name": "FreeRegistration",
                "help": "Free registration allowed",
                "json": "free_registration",
                "type": "bool"
            },
            {
                "name": "MaxUploadMb",
                "help": "Maximum size of user's upload",
                "json": "max_upload_mb",
                "type": "int"
            },
            {
                "name": "MaxLinkedMessages",
                "help": "Maximum number of forwarded messages",
                "json": "max_linked_messages",
                "type": "int"
            },
            {
                "name": "MaxUsernamePartLength",
                "help": "Maximum chars for: family_name, given_name, patronymic if any",
                "json": "max_username_part_length",
                "type": "int"
            },
            {
                "name": "MaxGroupTitleLength",
                "help": "Maximum chars for group chat name",
                "json": "max_group_title_length",
                "type": "int"
            },
            {
                "name": "MaxRoleLength",
                "help": "Maximum chars for role in team",
                "json": "max_role_length",
                "type": "int"
            },
            {
                "name": "MaxMoodLength",
                "help": "Maximum chars for mood in team",
                "json": "max_mood_length",
                "type": "int"
            },
            {
                "name": "MaxMessageLength",
                "help": "Maximum chars for text message",
                "json": "max_message_length",
                "type": "int"
            },
            {
                "name": "MaxSectionLength",
                "help": "Maximum length for project and contact's sections names",
                "json": "max_section_length",
                "type": "int"
            },
            {
                "name": "MaxTagLength",
                "help": "Maximum length for tags",
                "json": "max_tag_length",
                "type": "int"
            },
            {
                "name": "MaxTaskTitleLength",
                "help": "Maximum length for task title",
                "json": "max_task_title_length",
                "type": "int"
            },
            {
                "name": "MaxTeams",
                "help": "Maximum teams for one account",
                "json": "max_teams",
                "type": "int"
            },
            {
                "name": "AfkAge",
                "help": "Max inactivity seconds",
                "json": "afk_age",
                "type": "int"
            },
            {
                "name": "AuthByPassword",
                "help": "Password authentication enabled",
                "json": "auth_by_password",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "AuthByQrCode",
                "help": "QR-code / link authentication enabled",
                "json": "auth_by_qr_code",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "AuthBySms",
                "help": "SMS authentication enabled",
                "json": "auth_by_sms",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "ICEServers",
                "help": "ICE servers for WebRTC",
                "json": "ice_servers",
                "type": "[]ICEServer"
            },
            {
                "name": "CustomServer",
                "help": "True for onpremise installation",
                "json": "custom_server",
                "type": "bool"
            },
            {
                "name": "InstallationType",
                "help": "name of instalation",
                "json": "installation_type",
                "type": "string"
            },
            {
                "name": "IsTesting",
                "help": "testing installation",
                "json": "is_testing",
                "type": "bool"
            },
            {
                "name": "Metrika",
                "help": "yandex metrika counter id",
                "json": "metrika",
                "type": "string"
            },
            {
                "name": "MinSearchLength",
                "help": "minimal chars number for starting global search",
                "json": "min_search_length",
                "type": "int"
            },
            {
                "name": "ResendTimeout",
                "help": "resend message in n seconds if no confirmation from server given",
                "json": "resend_timeout",
                "type": "int"
            },
            {
                "name": "SentryDsnJS",
                "help": "frontent sentry.io settings",
                "json": "sentry_dsn_js",
                "type": "string"
            },
            {
                "name": "ServerDrafts",
                "help": "message drafts saved on server",
                "json": "server_drafts",
                "type": "bool"
            },
            {
                "name": "FirebaseAppId",
                "help": "Firebase application id for web-push notifacations",
                "json": "firebase_app_id",
                "type": "string"
            },
            {
                "name": "FirebaseSenderId",
                "help": "Firebase sender id for web-push notifacations",
                "json": "firebase_sender_id",
                "type": "string"
            },
            {
                "name": "Calls",
                "help": "Calls functions enabled",
                "json": "calls",
                "type": "bool"
            },
            {
                "name": "MobileCalls",
                "help": "Calls functions enabled for mobile applications",
                "json": "mobile_calls",
                "type": "bool"
            },
            {
                "name": "CallsRecord",
                "help": "Calls record enabled",
                "json": "calls_record",
                "type": "bool"
            },
            {
                "name": "OnlyOneDevicePerCall",
                "help": "Disallow call from multiply devices. Experimental",
                "json": "only_one_device_per_call",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "MaxParticipantsPerCall",
                "help": "Maximum number of participants per call",
                "json": "max_participants_per_call",
                "type": "int",
                "omitempty": true
            },
            {
                "name": "SafariPushId",
                "help": "Safari push id for web-push notifacations",
                "json": "safari_push_id",
                "type": "string"
            },
            {
                "name": "Terms",
                "help": "Team entity naming. Experimental.",
                "json": "terms",
                "type": "Terms"
            },
            {
                "name": "SingleGroupTeams",
                "help": "Cross team communication. Experimental.",
                "json": "single_group_teams",
                "type": "bool"
            },
            {
                "name": "WikiPages",
                "help": "Wiki pages in chats. Experimental",
                "json": "wiki_pages",
                "type": "bool"
            },
            {
                "name": "AllowAdminMute",
                "help": "Wiki pages in chats. Experimental",
                "json": "allow_admin_mute",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "TaskChecklist",
                "help": "Deprecated. Always true",
                "json": "task_checklist",
                "type": "bool"
            },
            {
                "name": "ReadonlyGroups",
                "help": "Deprecated. Always true",
                "json": "readonly_groups",
                "type": "bool"
            },
            {
                "name": "TaskDashboard",
                "help": "Deprecated. Always true",
                "json": "task_dashboard",
                "type": "bool"
            },
            {
                "name": "TaskMessages",
                "help": "Deprecated. Always true",
                "json": "task_messages",
                "type": "bool"
            },
            {
                "name": "TaskPublic",
                "help": "Deprecated. Always true",
                "json": "task_public",
                "type": "bool"
            },
            {
                "name": "TaskTags",
                "help": "Deprecated. Always true",
                "json": "task_tags",
                "type": "bool"
            }
        ]
    },
    {
        "name": "MessageLinkPreview",
        "help": "Website title and description",
        "fields": [
            {
                "name": "Title",
                "help": "",
                "json": "title",
                "type": "string"
            },
            {
                "name": "Description",
                "help": "",
                "json": "description",
                "type": "string",
                "omitempty": true
            }
        ]
    },
    {
        "name": "MessageLink",
        "help": "Checked message links. In short: \"Click here: {link.Pattern}\" => \"Click here: <a href='{link.Url}'>{link.Text}</a>\"",
        "fields": [
            {
                "name": "Pattern",
                "help": "Text fragment that should be replaced by link",
                "json": "pattern",
                "type": "string"
            },
            {
                "name": "Url",
                "help": "Internal (tadateam://) or external link",
                "json": "url",
                "type": "string"
            },
            {
                "name": "Text",
                "help": "Text replacement.",
                "json": "text",
                "type": "string"
            },
            {
                "name": "Preview",
                "help": "Optional preview info, for websites",
                "json": "preview",
                "type": "MessageLinkPreview",
                "null": true,
                "omitempty": true
            },
            {
                "name": "Uploads",
                "help": "Optional upload info",
                "json": "uploads",
                "type": "[]Upload",
                "omitempty": true
            },
            {
                "name": "NoPreview",
                "help": "Website previews disabled",
                "json": "nopreview",
                "type": "bool",
                "omitempty": true
            },
            {
                "name": "YoutubeId",
                "help": "Optional youtube movie id",
                "json": "youtube_id",
                "type": "string",
                "omitempty": true
            }
        ]
    }
]

```
