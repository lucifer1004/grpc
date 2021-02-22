fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::configure()
        .build_server(true)
        .compile(
            &[
                "src/protos/puppet.proto",
                "src/protos/puppet/base.proto",
                "src/protos/puppet/contact.proto",
                "src/protos/puppet/event.proto",
                "src/protos/puppet/file_box.proto",
                "src/protos/puppet/friendship.proto",
                "src/protos/puppet/message.proto",
                "src/protos/puppet/room.proto",
                "src/protos/puppet/room_invitation.proto",
                "src/protos/puppet/room_member.proto",
                "src/protos/puppet/tag.proto",
            ],
            &["src/protos"],
        )?;
    Ok(())
}