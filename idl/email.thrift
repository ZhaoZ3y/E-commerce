namespace go email

struct EmailReq{
    1: string from,
    2: string to,
    3: string content_type,
    4: string subject,
    5: string content,
}

struct EmailResp{
}

service EmailService {
    EmailResp Send(1: EmailReq req),
}