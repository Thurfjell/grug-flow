const std = @import("std");

const net = std.Io.net;

fn handleClient(io: std.Io, mut_conn: net.Stream) void {
    var conn = mut_conn;
    defer conn.close(io);

    var recv_buffer: [1024]u8 = undefined;
    var send_buffer: [1024]u8 = undefined;

    var reader = conn.reader(io, &recv_buffer);
    var writer = conn.writer(io, &send_buffer);

    var http_server = std.http.Server.init(&reader.interface, &writer.interface);

    var req = http_server.receiveHead() catch return;
    std.log.info("request path: {s}", .{req.head.target});

    req.respond("hello from concurrent task\n", .{
        .extra_headers = &.{
            .{
                .name = "content-type",
                .value = "text/plain",
            },
        },
    }) catch return;
}

pub fn main(init: std.process.Init) !void {
    const io = init.io;

    const addr = try net.IpAddress.parse("127.0.0.1", 8080);

    var server = try addr.listen(io, .{});
    defer server.deinit(io);

    std.log.info("server listening at http://127.0.0.1:8080", .{});

    var group: std.Io.Group = .init;
    defer group.cancel(io);

    while (true) {
        const conn = try server.accept(io);

        group.async(io, handleClient, .{ io, conn });
    }
}
