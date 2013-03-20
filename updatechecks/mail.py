import smtplib
import string

def send_mail(to_addr, from_addr, subject, body, host="localhost"):
""" Send an email. """
    BODY = string.join((
            "From: %s" % from_addr,
            "To: %s" % to_addr,
            "Subject: %s" % subject,
            "",
            body
            ), "\r\n")
    server = smtplib.SMTP(host)
    server.sendmail(from_addr, [TO_addr], BODY)
    server.quit()
