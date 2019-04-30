FROM prismagraphql/prisma:1.31
COPY ./prisma-patch.sh /app/start.sh
CMD [ "/app/start.sh" ]
